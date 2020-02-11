package paypal

import "github.com/missionpay/gateway-switch/app/interfaces/customer"

var transaction *Transaction

func Initialize(trans *Transaction) {
	transaction = trans
}

func Tokenize() (status customer.GatewayStatus) {
	return
}

func Authorize() (status customer.GatewayStatus) {
	return
}

func Capture() (status customer.GatewayStatus) {
	return
}

func Void() (status customer.GatewayStatus) {
	return
}

func Refund() (status customer.GatewayStatus) {
	return
}

func Webhook() (status customer.GatewayStatus) {
	return
}
