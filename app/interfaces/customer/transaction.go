package customer

type Transaction interface {
	AddDescription(description string)
	GetDescription() (description string)

	AddInvoiceNumber(invoiceNumber string)
	GetInvoiceNumber() (invoiceNumber string)

	SetGateway(gateway Gateway)
	GetGateway() (gateway Gateway)
}
