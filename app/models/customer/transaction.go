package customer

import (
	"github.com/missionpay/app/interfaces/customer"
	"time"
)

type Transaction struct {
	TransactionID string
	Description   string
	CreateDate    time.Time
	ModifiedDate  time.Time
	Gateway customer.Gateway
	Cart          Cart
}

func (trans *Transaction) AddDescription(description string) {
	trans.Description = description
}

func (trans *Transaction) GetDescription() (description string) {
	return trans.Description
}

func (trans *Transaction) SetGateway(gateway customer.Gateway) {
	trans.Gateway = gateway
}

func (trans *Transaction) GetGateway() (gateway customer.Gateway) {
	return trans.Gateway
}
