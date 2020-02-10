package customer

type Customer interface {
	SetFirstName(firstName string)
	GetFirstName() (firstName string)

	SetMiddleName(middleName string)
	GetMiddleName() (middleName string)

	SetLastName(lastName string)
	GetLastName() (lastName string)

	SetPhone(phone string)
	GetPhone() (phone string)

	SetEmail(email string)
	GetEmail() (email string)

	SetShippingAddress(shippingAddress Address)
	GetShippingAddress() (shippingAddress Address)

	AddCustomField(key, value string)
	GetCustomField(key string) (value string)

	NewTransaction()
	GetTransactionByID(transactionID string) (transaction Transaction)
	GetAllTransactions() (transactions []Transaction)

	Retrieve(customerID string) (err error)
	Save() (err error)
}
