package rest

import (
	"errors"
	"freemasonry.cc/blockchain/core"
	"freemasonry.cc/blockchain/util"
	types2 "freemasonry.cc/blockchain/x/chat/types"
	"freemasonry.cc/trerr"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"net/http"
)


func BroadcastTxHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := core.BuildLog("BroadcastTxHandlerFn", core.LmChainRest)

		var txBytes []byte
		if r.Body != nil {
			txBytes, _ = ioutil.ReadAll(r.Body)
		}

		txResponse := types2.BroadcastTxResponse{
			BaseResponse: types2.BaseResponse{Info: "", Status: 0},
			TxHash:       "",
			Height:       0,
		}
		tx, _ := clientCtx.TxConfig.TxDecoder()(txBytes) 
		stdTx, err := txToStdTx(clientCtx, tx)           
		if err != nil {
			txResponse.Info = err.Error()
			SendReponse(w, clientCtx, txResponse)
			return
		}
		msgs := stdTx.GetMsgs() 
		fee := stdTx.Fee        
		memo := stdTx.Memo      

		log.Info("")

		
		err = broadcastMsgCheck(msgs, fee, memo)
		if err != nil {
			errmsg := trerr.TransError(err.Error())
			txResponse.Info = errmsg.Error()
			SendReponse(w, clientCtx, txResponse)
			return
		}

		res, err := clientCtx.BroadcastTx(txBytes)
		if err != nil {
			txResponse.Info = err.Error()
			SendReponse(w, clientCtx, txResponse)
			return
		}

		if res.Code == 0 { 
			txResponse.Status = 1
		} else {
			txResponse.Status = 0
		}

		txResponse.Code = res.Code
		txResponse.Codespace = res.Codespace
		txResponse.TxHash = res.TxHash
		txResponse.Info = parseErrorCode(res.Code, res.Codespace, res.RawLog)
		txResponse.Height = res.Height
		SendReponse(w, clientCtx, txResponse)
	}
}


func parseErrorCode(code uint32, codeSpace string, rowlog string) string {
	if codeSpace == sdkErrors.RootCodespace {
		if code == sdkErrors.ErrInsufficientFee.ABCICode() { 
			return FeeIsTooLess
		} else if code == sdkErrors.ErrOutOfGas.ABCICode() { 
			return ErrorGasOut
		} else if code == sdkErrors.ErrUnauthorized.ABCICode() { 
			return ErrUnauthorized
		} else if code == sdkErrors.ErrWrongSequence.ABCICode() { 
			return ErrWrongSequence
		}
	}
	return rowlog
}

func broadcastMsgCheck(msgs []sdk.Msg, fee legacytx.StdFee, memo string) (err error) {
	for _, msg := range msgs {
		msgType := proto.MessageName(msg)
		if txHandles.HaveRegistered(msgType) { 
			msgByte, err := util.Json.Marshal(msg)
			if err != nil {
				return err
			}
			err = txHandles.Handle(msgType, msgByte, fee, memo)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func txToStdTx(clientCtx client.Context, tx sdk.Tx) (*legacytx.StdTx, error) {
	signingTx, ok := tx.(signing.Tx)
	if !ok {
		return nil, errors.New("tx to stdtx error")
	}
	stdTx, err := clienttx.ConvertTxToStdTx(clientCtx.LegacyAmino, signingTx)
	if err != nil {
		return nil, err
	}
	return &stdTx, nil
}
