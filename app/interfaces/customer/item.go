package customer

type Item interface {
	GetSku() (sku string)

	SetName(name string)
	GetName() (name string)

	SetDescription(description string)
	GetDescription() (description string)

	SetPrice(price int)
	GetPrice() (price int)

	SetQuantity(quantity int)
	GetQuantity() (quantity int)
}
