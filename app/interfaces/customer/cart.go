package customer

type Cart interface {
	IsPaid() bool
	GetCartID() string
	GetTotalDue() int
	GetTotalPaid() int

	SetEncryptedCard(encyptedCard string)
	GetEncryptedCard() (encryptedCard string)

	SetBillingSameAsShippingAddress(isSame bool)
	GetIsBillingSameAsShippingAddress() (isSame bool)

	AddTax(tax int)
	GetTax() (tax int)
	AddShipping(shipping int)
	GetShipping() (shipping int)
	AddHandlingFee(handlingFee int)
	GetHandlingFee() (handlingFee int)
	AddShippingDiscount(shippingDiscount int)
	GetShippingDiscount() (shippingDiscount int)
	AddInsurance(insurance int)
	GetInsurance() (insurance int)

	AddItem(item Item)
	AddItems(items ...Item)
	UpdateItemQuantity(item Item, quantity int)
	ClearAllItems()
	GetItems() (items []Item)
	GetItemsTotalQuantity() (total int)
	GetItemsTotalDue() (total int)
}
