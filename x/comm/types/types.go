package types

const (
	MSG_GATEWAY_REGISTER  = "comm/MsgGatewayRegister"
	MSG_ADDRESS_BOOK_SAVE = "comm/MsgAddressBookSave"
)


type Gateway struct {
	
	GatewayAddress string `json:"gateway_address"`
	
	GatewayName string `json:"gateway_name"`
	
	GatewayNum []GatewayNumIndex `json:"gateway_num"`
}


type GatewayNumIndex struct {
	
	GatewayAddress string `json:"gateway_address"`
	
	NumberIndex string `json:"number_index"`
	
	NumberEnd []string `json:"number_end"`
	
	Status int64 `json:"status"`
	
	Validity int64 `json:"validity"`
}
