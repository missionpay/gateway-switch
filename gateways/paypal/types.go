package paypal

type Payer struct {
	PaymentMethod string `json:"payment_method"`
}

type Amount struct {
	Total    string `json:"total"`
	Currency string `json:"currency"`
	Details  struct {
		Subtotal         string `json:"subtotal"`
		Tax              string `json:"tax"`
		Shipping         string `json:"shipping"`
		HandlingFee      string `json:"handling_fee"`
		ShippingDiscount string `json:"shipping_discount"`
		Insurance        string `json:"insurance"`
	} `json:"details"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Tax         string `json:"tax"`
	Sku         string `json:"sku"`
	Currency    string `json:"currency"`
}

type ShippingAddress stuct {
	RecipientName string `json:"recipient_name"`
	Line1         string `json:"line1"`
	Line2         string `json:"line2"`
	City          string `json:"city"`
	CountryCode   string `json:"country_code"`
	PostalCode    string `json:"postal_code"`
	Phone         string `json:"phone"`
	State         string `json:"state"`
}

type ItemList struct {
	Items []Item `json:"items"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method"`
}

type Transaction struct {
		Description    string `json:"description"`
		Custom         string `json:"custom"`
		InvoiceNumber  string `json:"invoice_number"`
		PaymentOptions PaymentOptions `json:"payment_options"`
		SoftDescriptor string `json:"soft_descriptor"`
		ItemList   ItemList `json:"item_list"`
}

type RedirectUrls struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type TokenRequest struct {
		Intent string `json:"intent"`
		Payer Payer `json:"payer"`
		NoteToPayer  string `json:"note_to_payer"`
		Transactions []Transaction `json:"transactions"`
		RedirectUrls RedirectUrls `json:"redirect_urls"`
}