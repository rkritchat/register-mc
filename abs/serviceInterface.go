package abs

type ServiceInterface interface {
	ValidateRequestMsg() error
	ValidateBusinessRule() error
	Execute() (string, error)
}
