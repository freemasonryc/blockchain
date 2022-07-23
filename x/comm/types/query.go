package types

const (
	QueryGatewayInfo = "gateway_info"
	QueryGatewayNum  = "gateway_num"
)


type QueryGatewayInfoParams struct {
	GatewayAddress  string `json:"gateway_address"`
	GatewayNumIndex string `json:"gateway_num_index"`
}
