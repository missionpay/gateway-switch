package customer

type Gateway interface {
	Tokenize()
	Authorize()
	Capture()
	Void()
	Refund()
	Webhook()
}
