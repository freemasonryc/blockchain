package client

import (
	"bytes"
	"errors"
	"freemasonry.cc/blockchain/util"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"net/url"
	"strconv"
	"strings"
	"time"

	
	"io/ioutil"
	"net/http"
)

type StringEvent struct {
	Type       string      `json:"type,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type ABCIMessageLog struct {
	MsgIndex uint16 `json:"msg_index"`
	Log      string `json:"log"`

	Events []StringEvent `json:"events"`
}

type TxDetail struct {
	Height string `json:"height"`
	Status string `json:"status"`
	Txhash string `json:"txhash"`
	Error  string `json:"error"`
}


func GetRequest(server, params string) (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, 
			MaxIdleConnsPerHost: 512,  
		},
		Timeout: time.Second * 60, 
	}
	bodyReader := strings.NewReader("")
	req, err := http.NewRequest("GET", server+params, bodyReader)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}
	
	return string(body), err
}

func PostValuesRequest(server, url string, values url.Values) (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, 
			MaxIdleConnsPerHost: 512,  
		},
		Timeout: time.Second * 60, 
	}
	req, err := http.NewRequest("POST", server+url, nil)
	req.PostForm = values
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	
	if resp.StatusCode != 200 {
		if len(body) == 0 {
			return "", errors.New("err code return:" + strconv.Itoa(resp.StatusCode))
		}
		return "", errors.New(string(body))
	}
	
	return string(body), err
}

func PostRequest(server, url string, params []byte) (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, 
			MaxIdleConnsPerHost: 512,  
		},
		Timeout: time.Second * 60, 
	}
	req, err := http.NewRequest("POST", server+url, bytes.NewReader(params))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		if len(body) == 0 {
			return "", errors.New("error code:" + strconv.Itoa(resp.StatusCode))
		}
		return "", errors.New(string(body))
	}

	return string(body), err
}

func PostRequestByTimeout(server, url string, params []byte, timeout time.Duration) (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, 
			MaxIdleConnsPerHost: 512,  
		},
		Timeout: timeout, 
	}
	req, err := http.NewRequest("POST", server+url, bytes.NewReader(params))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}
	
	return string(body), err
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


func unmarshalMsg(msg sdk.Msg, obj interface{}) error {
	msgByte, err := util.Json.Marshal(msg)
	if err != nil {
		return err
	}
	return util.Json.Unmarshal(msgByte, &obj)
}
