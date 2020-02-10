package customer

import "github.com/missionpay/app/models/customer"

type Cart interface {
	IsPaid() bool
	GetCartID() string
	GetTotalDue() int
	GetTotalPaid() int

	SetEncryptedCard(encyptedCard string)
	GetEncryptedCard() (encryptedCard string)

	SetBillingSameAsShippingAddress(isSame bool)
	GetIsBillingSameAsShippingAddress() (isSame bool)

	AddItem(item customer.Item)
	AddItems(items ...customer.Item)
	UpdateItemQuantity(item customer.Item, quantity int)
	ClearAllItems()
	GetItems() (items []customer.Item)
	GetItemsTotalQuantity() (total int)
	GetItemsTotalDue() (total int)
}
