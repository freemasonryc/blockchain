package rest

var (
	FeeIsTooLess     = "fee is too less"
	ErrorGasOut      = "The gas consumed exceeds the upper limit set by the client"
	ErrWrongSequence = "account serial number expired, the reason may be: node block behind or repeatedly sent messages"
	ErrUnauthorized  = "signature verification failed, invalid chainid or account number"
)
