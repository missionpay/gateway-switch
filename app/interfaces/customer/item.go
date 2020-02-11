package customer

type Item interface {
	SetSku(sku string)
	GetSku() (sku string)

	SetName(name string)
	GetName() (name string)

	SetDescription(description string)
	GetDescription() (description string)

	SetPrice(price int)
	GetPrice() (price int)

	SetTax(tax int)
	GetTax() (tax int)

	SetCurrency(currency string)
	GetCurrency() (currency string)

	SetQuantity(quantity int)
	GetQuantity() (quantity int)
}
