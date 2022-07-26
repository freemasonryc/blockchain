package address

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sort"

	"github.com/cosmos/cosmos-sdk/internal/conv"
	"github.com/cosmos/cosmos-sdk/types/errors"
)


const Len = sha256.Size


type Addressable interface {
	Address() []byte
}


func Hash(typ string, key []byte) []byte {
	hasher := sha256.New()
	_, err := hasher.Write(conv.UnsafeStrToBytes(typ))

	errors.AssertNil(err)
	th := hasher.Sum(nil)

	hasher.Reset()
	_, err = hasher.Write(th)
	errors.AssertNil(err)
	_, err = hasher.Write(key)
	errors.AssertNil(err)
	return hasher.Sum(nil)
}


func Compose(typ string, subAddresses []Addressable) ([]byte, error) {
	as := make([][]byte, len(subAddresses))
	totalLen := 0
	var err error
	for i := range subAddresses {
		a := subAddresses[i].Address()
		as[i], err = LengthPrefix(a)
		if err != nil {
			return nil, fmt.Errorf("not compatible sub-adddress=%v at index=%d [%w]", a, i, err)
		}
		totalLen += len(as[i])
	}

	sort.Slice(as, func(i, j int) bool { return bytes.Compare(as[i], as[j]) <= 0 })
	key := make([]byte, totalLen)
	offset := 0
	for i := range as {
		copy(key[offset:], as[i])
		offset += len(as[i])
	}
	return Hash(typ, key), nil
}



func Module(moduleName string, key []byte) []byte {
	mKey := append([]byte(moduleName), 0)
	return Hash("module", append(mKey, key...))
}


func Derive(address []byte, key []byte) []byte {
	return Hash(conv.UnsafeBytesToStr(address), key)
}
