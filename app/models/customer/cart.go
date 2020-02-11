package customer

type Cart struct {
	CartID                         string
	TotalDue                       int
	TotalPaid                      int
	EncryptedCard                  string
	IsBillingSameAsShippingAddress bool
	BillingAddress                 Address // billing address on cart for recurring purposes
	Items                          []Item
}

func (cart *Cart) IsPaid() bool {
	return cart.TotalPaid > cart.TotalPaid
}

func (cart *Cart) GetCartID() string {
	return cart.CartID
}

func (cart *Cart) GetTotalDue() int {
	return cart.TotalDue
}

func (cart *Cart) GetTotalPaid() int {
	return cart.TotalPaid
}

func (cart *Cart) SetEncryptedCard(encryptedCard string) {
	cart.EncryptedCard = encryptedCard
}

func (cart *Cart) GetEncryptedCard() (encryptedCard string) {
	return cart.EncryptedCard
}

func (cart *Cart) SetBillingSameAsShippingAddress(isSame bool) {
	cart.IsBillingSameAsShippingAddress = isSame
}

func (cart *Cart) GetIsBillingSameAsShippingAddress() (isSame bool) {
	return cart.IsBillingSameAsShippingAddress
}

func (cart *Cart) AddItem(item Item) {
	cart.Items = append(cart.Items, item)
}

func (cart *Cart) AddItems(items ...Item) {
	for _, item := range items {
		cart.Items = append(cart.Items, item)
	}
}

func (cart *Cart) UpdateItemQuantity(item Item, quantity int) {
	for i, cartItem := range cart.Items {
		if item.Sku == cartItem.Sku || item.Name == cartItem.Name {
			if quantity > 0 {
				cart.Items[i].Quantity = quantity
			} else {
				// Remove Item from array if quantity is Zero
				copy(cart.Items[i:], cart.Items[i+1:])
			}
		}
	}
}

func (cart *Cart) ClearAllItems() {
	cart.Items = []Item{}
}

func (cart *Cart) GetItems() (items []Item) {
	return cart.Items
}

func (cart *Cart) GetItemsTotalQuantity() (total int) {
	for _, item := range cart.Items {
		total += item.Quantity
	}

	return
}

func (cart *Cart) GetItemsTotalDue() (total int) {
	for _, item := range cart.Items {
		total += item.Price
	}

	return
}
