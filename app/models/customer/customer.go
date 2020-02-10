package customer

import (
	"os"
	"time"

	"github.com/catmullet/dynamo"
	"github.com/google/uuid"
	"github.com/missionpay/app/interfaces/customer"
)

type Customer struct {
	CustomerID      string
	FirstName       string
	MiddleName      string
	LastName        string
	Phone           string
	Email           string
	ShippingAddress customer.Address
	CustomFields    map[string]string
	Transactions    map[string]Transaction
}

func (cust *Customer) SetFirstName(firstName string) {
	cust.FirstName = firstName
}

func (cust *Customer) GetFirstName() (firstName string) {
	return cust.FirstName
}

func (cust *Customer) SetMiddleName(middleName string) {
	cust.MiddleName = middleName
}

func (cust *Customer) GetMiddleName() (middleName string) {
	return cust.MiddleName
}

func (cust *Customer) SetLastName(lastName string) {
	cust.LastName = lastName
}

func (cust *Customer) GetLastName() (lastName string) {
	return cust.LastName
}

func (cust *Customer) SetPhone(phone string) {
	cust.Phone = phone
}

func (cust *Customer) GetPhone() (phone string) {
	return cust.Phone
}

func (cust *Customer) SetEmail(email string) {
	cust.Email = email
}

func (cust *Customer) GetEmail() (email string) {
	return cust.Email
}

func (cust *Customer) SetShippingAddress(shippingAddress Address) {
	cust.ShippingAddress = shippingAddress
}

func (cust *Customer) GetShippingAddress() (shippingAddress Address) {
	return cust.ShippingAddress
}

func (cust *Customer) AddCustomField(key, value string) {
	if cust.CustomFields != nil {
		cust.CustomFields[key] = value
	} else {
		cust.CustomFields = make(map[string]string)
	}
}

func (cust *Customer) GetCustomField(key string) (value string) {
	if val, ok := cust.CustomFields[key]; ok {
		value = val
	}

	return
}

func (cust *Customer) NewTransaction() (newTransaction Transaction) {
	id := uuid.New()
	newTransaction.TransactionID = id.String()

	if cust.Transactions != nil {
		cust.Transactions[id.String()] = newTransaction
	} else {
		cust.Transactions = make(map[string]Transaction)
		cust.Transactions[id.String()] = newTransaction
	}

	return
}

func (cust *Customer) GetTransactionByID(transactionID string) (transaction Transaction) {
	if val, ok := cust.Transactions[transactionID]; ok {
		transaction = val
	}

	return
}

func (cust *Customer) GetAllTransactions() (transactions map[string]Transaction) {
	if cust.Transactions != nil {
		transactions = cust.Transactions
	} else {
		cust.Transactions = make(map[string]Transaction)
		transactions = cust.Transactions
	}

	return
}

func (cust *Customer) Retrieve(customerID string) (err error) {
	err = dynamo.GetItem(os.Getenv("MISSIONPAY_TRANSACTIONS_TABLE"), map[string]interface{}{"CustomerID": customerID}, &cust)
	return
}

func (cust *Customer) Save() (err error) {
	err = dynamo.Put(os.Getenv("MISSIONPAY_TRANSACTIONS_TABLE"), cust, 17520*time.Hour) // 2 Year delete if no activity
	return
}
