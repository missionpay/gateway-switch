package customer

type Gateway interface {
	
	Initialize()

	Tokenize()
	Authorize()
	Capture()
	Void()
	Refund()
	Webhook()
}