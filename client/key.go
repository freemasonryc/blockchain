package client

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"freemasonry.cc/blockchain/cmd/config"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	cmssecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/go-bip39"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tyler-smith/go-bip39"
)

func NewSecretKey() *SecretKey {
	return &SecretKey{}
}


type SecretKey struct {
}


func (k *SecretKey) CreateSeedWord() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}


	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}


func (k *SecretKey) CreateAccountFromSeed(seed string) (*CosmosWallet, error) {
	seed, err := genSeed(seed)
	if err != nil {
		
		return nil, err
	}
	return genAddress(seed, 0)
}

func (k *SecretKey) CreateAccount(seed string, index int) (*CosmosWallet, error) {
	seed, err := genSeed(seed)
	if err != nil {
		
		return nil, err
	}
	return genAddress(seed, index)
}


func (k *SecretKey) CreateAccountFromPriv(priv string) (*CosmosWallet, error) {
	var privkey secp256k1.PrivKey
	privKeyBytes, err := hex.DecodeString(priv)
	if err != nil {
		return nil, err
	}
	privkey = privKeyBytes
	pubKey := privkey.PubKey()
	bech32Addr, err := bech32.ConvertAndEncode(config.Bech32Prefix, pubKey.Address())
	if err != nil {
		return nil, err
	}
	return &CosmosWallet{PrivateKey: priv, PublicKey: hex.EncodeToString(pubKey.Bytes()), Address: bech32Addr}, nil
}

func (k *SecretKey) Sign(addr *CosmosWallet, msg []byte) ([]byte, error) {
	var privkey secp256k1.PrivKey
	privKeyBytes, err := hex.DecodeString(addr.PrivateKey)
	if err != nil {
		return nil, err
	}
	privkey = privKeyBytes
	privkey1 := cmssecp256k1.PrivKey{Key: privkey}
	return privkey1.Sign(msg)
}

func (k *SecretKey) SignString(privateKey string, msg []byte) ([]byte, error) {
	var privkey secp256k1.PrivKey
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	privkey = privKeyBytes
	privkey1 := cmssecp256k1.PrivKey{Key: privkey}
	return privkey1.Sign(msg)
}

func (k *SecretKey) CheckSign(publicKeyHex string, msg, sign []byte) (bool, error) {
	pubkeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, err
	}
	pubKey := cmssecp256k1.PubKey{Key: pubkeyBytes}
	return pubKey.VerifySignature(msg, sign), nil
}

const PubKeyHashAddrID = 0x76
const PrivateKeyID = 0x80
const CoinType = 118

func genSeed(seedWords string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(seedWords, "")
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(seed), nil
}

func genAddress(seed string, index int) (*CosmosWallet, error) {
	pkb, err := hex.DecodeString(seed)
	if err != nil {
		return nil, err
	}
	return bip44(pkb, index)
}

func bip44(pkb []byte, index int) (*CosmosWallet, error) {
	net := chaincfg.MainNetParams

	altcointype := CoinType
	net.PubKeyHashAddrID = byte(PubKeyHashAddrID)
	net.PrivateKeyID = byte(PrivateKeyID)

	ext, err := hdkeychain.NewMaster(pkb, &net)
	if err != nil {
		return nil, errors.New("Bip32 root key generation failed:" + err.Error())
	}







	purpose, err := ext.Derive(44 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, errors.New("Bip44 purpose (m/44') failed:" + err.Error())
	}


	coinType, err := purpose.Derive(uint32(altcointype) + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, errors.New("Bip44 coin_type (m/44'/coin_type') failed:" + err.Error())
	}


	acct0, err := coinType.Derive(0 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return nil, errors.New("Bip44 first account (m/44'/coin_type'/0') failed:" + err.Error())
	}

	acct0External, err := acct0.Derive(0)
	if err != nil {
		return nil, err
	}

	receive, err := acct0External.Derive(uint32(index))
	if err != nil {
		return nil, errors.New("Failed to create receive address:" + err.Error())
	}

	privk, err := receive.ECPrivKey()
	if err != nil {
		return nil, errors.New("ECPrivKey() failed:" + err.Error())
	}


	pubk, err := receive.ECPubKey()
	if err != nil {
		return nil, errors.New("ECPubKey() failed:" + err.Error())
	}

	var addr CosmosWallet

	privkSer := privk.Serialize()

	

	var privkey secp256k1.PrivKey = privk.Serialize()
	var pubkey secp256k1.PubKey = privkey.PubKey().(secp256k1.PubKey)

	bech32Addr, err := bech32.ConvertAndEncode(config.Bech32Prefix, pubkey.Address())
	if err != nil {
		return nil, errors.New("bech32.ConvertAndEncode() failed:" + err.Error())
	}

	privkSerHex := hex.EncodeToString(privkSer)
	pubkSerHex := hex.EncodeToString(pubk.SerializeCompressed())

	addr = CosmosWallet{bech32Addr, pubkSerHex, privkSerHex}

	return &addr, nil
}


type CosmosWallet struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publickey"`
	PrivateKey string `json:"privatekey"`
}

func (this *CosmosWallet) MarshalJson() []byte {
	data, _ := json.Marshal(this)
	return data
}
