package customer

type Item struct {
	Sku         string
	Name        string
	Description string
	Price       int
	Quantity    int
}

func (item *Item) GetSku() (sku string) {
	return item.Sku
}

func (item *Item) SetName(name string) {
	item.Name = name
}

func (item *Item) GetName() (name string) {
	return item.Name
}

func (item *Item) SetDescription(description string) {
	item.Description = description
}

func (item *Item) GetDescription() (description string) {
	return item.Description
}

func (item *Item) SetPrice(price int) {
	item.Price = price
}

func (item *Item) GetPrice() (price int) {
	return item.Price
}

func (item *Item) SetQuantity(quantity int) {
	item.Quantity = quantity
}

func (item *Item) GetQuantity() (quantity int) {
	return item.Quantity
}
