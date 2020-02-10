package customer

type Transaction interface {
	AddDescription(description string)
	GetDescription() (description string)

	SetGateway()
	GetGateway()
}
