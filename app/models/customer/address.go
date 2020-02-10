package customer

// Address default for billing and shipping addresses
type Address struct {
	Street1      string
	Street2      string
	City         string
	State        string
	Country      string
	Postal       string
	CustomFields map[string]string
}

// AddStreet1 Add first street line to address
func (addr *Address) AddStreet1(street1 string) {
	addr.Street1 = street1
}

// GetStreet1 Get first street line from address
func (addr *Address) GetStreet1() (street1 string) {
	return addr.Street1
}

func (addr *Address) AddStreet2(street2 string) {
	addr.Street2 = street2
}

func (addr *Address) GetStreet2() (street2 string) {
	return addr.Street2
}

func (addr *Address) AddCity(city string) {
	addr.City = city
}

func (addr *Address) GetCity() (city string) {
	return addr.City
}

func (addr *Address) AddState(state string) {
	addr.State = state
}

func (addr *Address) GetState() (state string) {
	return addr.State
}

func (addr *Address) AddCountry(country string) {
	addr.Country = country
}

func (addr *Address) GetCountry() (country string) {
	return addr.Country
}

func (addr *Address) AddPostal(postal string) {
	addr.Postal = postal
}

func (addr *Address) GetPostal() (postal string) {
	return addr.Postal
}

func (addr *Address) AddCustomField(key, value string) {
	if addr.CustomFields != nil {
		addr.CustomFields[key] = value
	} else {
		addr.CustomFields = make(map[string]string)
		addr.CustomFields[key] = value
	}
}

func (addr *Address) GetCustomField(key string) (value string) {
	if val, ok := addr.CustomFields[key]; ok {
		value = val
	}

	return
}

func (addr *Address) AddAddressByString(addressString string) (err error) {
	return
}

func Verify() (recommendation Address, isVerified bool, err error) {
	return
}
