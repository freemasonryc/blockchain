package types

const (
	MSG_GATEWAY_REGISTER     = "comm/MsgGatewayRegister"
	MSG_GATEWAY_DELEGATION   = "comm/MsgGatewayDelegation"
	MSG_GATEWAY_UNDELEGATION = "comm/MsgGatewayUndelegation"
)


type Gateway struct {
	
	GatewayAddress string `json:"gateway_address"`
	
	GatewayName string `json:"gateway_name"`
	
	GatewayUrl string `json:"gateway_url"`
	
	GatewayQuota int64 `json:"gateway_quota"`
	
	GatewayNum []GatewayNumIndex `json:"gateway_num"`
}


type GatewayNumIndex struct {
	
	GatewayAddress string `json:"gateway_address"`
	
	NumberIndex string `json:"number_index"`
	
	NumberEnd []string `json:"number_end"`
	
	Status int64 `json:"status"`
	
	Validity int64 `json:"validity"`
}


type ValidatorInfor struct {
	ValidatorConsAddr string `json:"validator_consaddr"` 
	ValidatorStatus   string `json:"validator_status"`   
	ValidatorPubAddr  string `json:"validator_pubaddr"`  
	ValidatorOperAddr string `json:"validator_operaddr"` 
	AccAddr           string `json:"acc_addr"`           
}
