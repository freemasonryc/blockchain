package signing

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)






type SignatureV2 struct {

	PubKey cryptotypes.PubKey



	Data SignatureData



	Sequence uint64
}




func SignatureDataToProto(data SignatureData) *SignatureDescriptor_Data {
	switch data := data.(type) {
	case *SingleSignatureData:
		return &SignatureDescriptor_Data{
			Sum: &SignatureDescriptor_Data_Single_{
				Single: &SignatureDescriptor_Data_Single{
					Mode:      data.SignMode,
					Signature: data.Signature,
				},
			},
		}
	case *MultiSignatureData:
		descDatas := make([]*SignatureDescriptor_Data, len(data.Signatures))

		for j, d := range data.Signatures {
			descDatas[j] = SignatureDataToProto(d)
		}

		return &SignatureDescriptor_Data{
			Sum: &SignatureDescriptor_Data_Multi_{
				Multi: &SignatureDescriptor_Data_Multi{
					Bitarray:   data.BitArray,
					Signatures: descDatas,
				},
			},
		}
	default:
		panic(fmt.Errorf("unexpected case %+v", data))
	}
}




func SignatureDataFromProto(descData *SignatureDescriptor_Data) SignatureData {
	switch descData := descData.Sum.(type) {
	case *SignatureDescriptor_Data_Single_:
		return &SingleSignatureData{
			SignMode:  descData.Single.Mode,
			Signature: descData.Single.Signature,
		}
	case *SignatureDescriptor_Data_Multi_:
		multi := descData.Multi
		datas := make([]SignatureData, len(multi.Signatures))

		for j, d := range multi.Signatures {
			datas[j] = SignatureDataFromProto(d)
		}

		return &MultiSignatureData{
			BitArray:   multi.Bitarray,
			Signatures: datas,
		}
	default:
		panic(fmt.Errorf("unexpected case %+v", descData))
	}
}

var _, _ codectypes.UnpackInterfacesMessage = &SignatureDescriptors{}, &SignatureDescriptor{}


func (sds *SignatureDescriptors) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, sig := range sds.Signatures {
		err := sig.UnpackInterfaces(unpacker)

		if err != nil {
			return err
		}
	}

	return nil
}


func (sd *SignatureDescriptor) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return unpacker.UnpackAny(sd.PublicKey, new(cryptotypes.PubKey))
}
