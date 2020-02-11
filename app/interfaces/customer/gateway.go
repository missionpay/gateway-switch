package customer

type Gateway interface {
	Initialize(transaction *Transaction)

	Tokenize() (status GatewayStatus)
	Authorize() (status GatewayStatus)
	Capture() (status GatewayStatus)
	Void() (status GatewayStatus)
	Refund() (status GatewayStatus)
	Webhook() (status GatewayStatus)
}

type GatewayStatus string

const Success GatewayStatus = "Success"
const Failed GatewayStatus = "Failed"
const Pending GatewayStatus = "Pending"
