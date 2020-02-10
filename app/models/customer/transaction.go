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
	Cart          Cart
}

func (trans *Transaction) AddDescription(description string) {
	trans.Description = description
}

func (trans *Transaction) GetDescription() (description string) {
	return trans.Description
}

func SetGateway(gateway customer.Gateway) {

}

func GetGateway() (gateway customer.Gateway)
