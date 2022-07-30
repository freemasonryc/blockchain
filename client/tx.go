package client

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"freemasonry.cc/blockchain/cmd/config"
	"freemasonry.cc/blockchain/core"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/shopspring/decimal"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	ttypes "github.com/tendermint/tendermint/types"
	evmhd "github.com/tharsis/ethermint/crypto/hd"
	"regexp"
	"strconv"
)

type TxInfo struct {
	Height             string   
	Txhash             string   
	Data               string   
	LastCommitHash     string   
	Datahash           string   
	ValidatorsHash     string   
	NextValidatorsHash string   
	ConsensusHash      string   
	Apphash            string   
	LastResultsHash    string   
	EvidenceHash       string   
	ProposerAddress    string   
	Txs                []string 
	Signatures         []Signature
}

type TxClient struct {
	ServerUrl string
	logPrefix string
}

func (this *TxClient) ConvertTxToStdTx(cosmosTx sdk.Tx) (*legacytx.StdTx, error) {
	signingTx, ok := cosmosTx.(xauthsigning.Tx)
	if !ok {
		return nil, errors.New("tx to stdtx error")
	}
	stdTx, err := tx.ConvertTxToStdTx(clientCtx.LegacyAmino, signingTx)
	if err != nil {
		return nil, err
	}
	return &stdTx, nil
}


func (this *TxClient) TermintTx2CosmosTx(signTxs ttypes.Tx) (sdk.Tx, error) {
	return clientCtx.TxConfig.TxDecoder()(signTxs)
}


func (this *TxClient) SignTx2Bytes(signTxs xauthsigning.Tx) ([]byte, error) {
	return clientCtx.TxConfig.TxEncoder()(signTxs)
}

func (this *TxClient) SetFee(signTxs xauthsigning.Tx) ([]byte, error) {
	return clientCtx.TxConfig.TxEncoder()(signTxs)
}


func (this *TxClient) FindByByte(txhash []byte) (resultTx *ctypes.ResultTx, notFound bool, err error) {
	notFound = false
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	node, err := clientCtx.GetNode()
	if err != nil {
		log.WithError(err).Error("GetNode")
		return
	}
	resultTx, err = node.Tx(context.Background(), txhash, true)
	if err != nil {
		
		notFound = this.isTxNotFoundError(err.Error())
		if notFound {
			err = nil
		} else {
			log.WithError(err).WithField("txhash", hex.EncodeToString(txhash)).Error("node.Tx")
		}
		return
	}
	return
}




func (this *TxClient) FindByHex(txhashStr string) (resultTx *ctypes.ResultTx, notFound bool, err error) {
	var txhash []byte
	notFound = false
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	txhash, err = hex.DecodeString(txhashStr)
	if err != nil {
		log.WithError(err).WithField("txhash", txhashStr).Error("hex.DecodeString")
		return
	}
	return this.FindByByte(txhash)
}


func (this *TxClient) isTxNotFoundError(errContent string) (ok bool) {
	errRegexp := `tx\ \([0-9A-Za-z]{64}\)\ not\ found`
	r, err := regexp.Compile(errRegexp)
	if err != nil {
		return false
	}
	if r.Match([]byte(errContent)) {
		return true
	} else {
		return false
	}
}


func (this *TxClient) SignAndSendMsg(address string, privateKey string, fee legacytx.StdFee, memo string, msg ...sdk.Msg) (txRes *core.BroadcastTxResponse, err error) {
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	
	seqDetail, err := this.FindAccountNumberSeq(address)
	if err != nil {
		return
	}

	
	signedTx, err := this.SignTx(privateKey, seqDetail, fee, memo, msg...)
	if err != nil {
		return
	}

	
	signPubkey, err := signedTx.GetPubKeys()
	if err != nil {
		return
	}

	signV2, _ := signedTx.GetSignaturesV2()
	senderAddrBytes := signV2[0].PubKey.Address().Bytes()
	signAddrBytes := signPubkey[0].Address().Bytes()
	
	

	if !bytes.Equal(signAddrBytes, senderAddrBytes) {
		return nil, errors.New("sign error")
	}
	

	
	signedTxBytes, err := this.SignTx2Bytes(signedTx)
	if err != nil {
		log.WithError(err).Error("SignTx2Bytes")
		return
	}
	
	
	txRes, err = this.Send(signedTxBytes)
	if txRes != nil {
		txRes.SignedTxStr = hex.EncodeToString(signedTxBytes)
	}
	return
}


func (this *TxClient) FindAccountNumberSeq(accountAddr string) (detail core.ChainAccountNumberSeqResponse, err error) {
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	reponse, err := GetRequest(this.ServerUrl, "/copyright/accountNumberSeq/"+accountAddr)
	if err != nil {
		log.WithError(err).Error("GetRequest")
		return
	}
	err = json.Unmarshal([]byte(reponse), &detail)
	if err != nil {
		log.WithError(err).Error("json.Unmarshal")
	}
	return
}


func (this *TxClient) Send(req []byte) (txRes *core.BroadcastTxResponse, err error) {
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	response, err := PostRequest(this.ServerUrl, "/chat/tx/broadcast", req)
	if err != nil {
		log.WithError(err).Error("PostRequest")
		return
	}
	txRes = &core.BroadcastTxResponse{}
	err = json.Unmarshal([]byte(response), txRes)
	if err != nil {
		log.WithError(err).Error("json.Unmarshal")
		return
	}
	return
}


func (this *TxClient) GasInfo(seqDetail core.ChainAccountNumberSeqResponse, msg ...sdk.Msg) (coin core.RealCoin, gas uint64, err error) {
	log := core.BuildLog(core.GetFuncName(), core.LmChainClient)
	/*seqDetail, err := this.FindAccountNumberSeq(msg.GetSigners()[0].String())
	  if err != nil {
	  	return
	  }*/

	
	
	
	

	clientFactory = clientFactory.WithSequence(seqDetail.Sequence)
	gasInfo, _, err := tx.CalculateGas(clientCtx, clientFactory, msg...)
	if err != nil {
		log.WithError(err).Error("tx.CalculateGas")
		return
	}
	gas = gasInfo.GasInfo.GasUsed * 2
	gasUsed := decimal.RequireFromString(strconv.Itoa(int(gas)))
	
	gasUsed = gasUsed.Mul(decimal.RequireFromString(core.MinimumGasPrices)).Add(decimal.NewFromFloat(1))
	gasDec, err := sdk.NewDecFromStr(gasUsed.StringFixed(6))
	if err != nil {
		log.WithError(err).Error("NewDecFromStr")
		return
	}
	gasDecCoin := sdk.NewDecCoinFromDec(config.BaseDenom, gasDec)
	amount := core.MustParseLedgerDecCoin(gasDecCoin)
	if decimal.RequireFromString(amount).LessThan(decimal.RequireFromString(core.ChainDefaultFeeStr)) {
		
		amount = core.ChainDefaultFeeStr
	}
	coin = core.NewRealCoinFromStr(config.DisplayDenom, amount)
	return
}


func (this *TxClient) SignTx(privateKey string, seqDetail core.ChainAccountNumberSeqResponse, fee legacytx.StdFee, memo string, msgs ...sdk.Msg) (xauthsigning.Tx, error) {
	log := core.BuildLog(core.GetStructFuncName(this), core.LmChainClient)
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		log.WithError(err).Error("hex.DecodeString")
		return nil, err
	}
	keyringAlgos := keyring.SigningAlgoList{evmhd.EthSecp256k1}
	algo, err := keyring.NewSigningAlgoFromString("eth_secp256k1", keyringAlgos)
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(privKeyBytes)
	
	if fee.Gas == flags.DefaultGasLimit {
		_, gas, err := this.GasInfo(seqDetail, msgs...)
		if err != nil {
			log.WithError(err).Error("CulGas")
			return nil, core.Errformat(err)
		}
		log.WithField("gas", gas).Info("CulGas:")
		fee.Gas = gas
	}
	signMode := clientCtx.TxConfig.SignModeHandler().DefaultMode()
	signerData := xauthsigning.SignerData{
		ChainID:       clientCtx.ChainID,
		AccountNumber: seqDetail.AccountNumber,
		Sequence:      seqDetail.Sequence,
	}
	txBuild, err := tx.BuildUnsignedTx(clientFactory, msgs...)
	if err != nil {
		log.WithError(err).Error("tx.BuildUnsignedTx")
		return nil, err
	}
	txBuild.SetGasLimit(fee.Gas)     
	txBuild.SetFeeAmount(fee.Amount) 
	txBuild.SetMemo(memo)            
	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   privKey.PubKey(),
		Data:     &sigData,
		Sequence: seqDetail.Sequence,
	}
	
	if err := txBuild.SetSignatures(sig); err != nil {
		log.WithError(err).Error("SetSignatures")
		return nil, err
	}
	signV2, err := tx.SignWithPrivKey(signMode, signerData, txBuild, privKey, clientCtx.TxConfig, seqDetail.Sequence)
	if err != nil {
		log.WithError(err).Error("SignWithPrivKey")
		return nil, err
	}
	err = txBuild.SetSignatures(signV2)
	if err != nil {
		log.WithError(err).Error("SetSignatures")
		return nil, err
	}

	signedTx := txBuild.GetTx()
	
	return signedTx, nil
}
