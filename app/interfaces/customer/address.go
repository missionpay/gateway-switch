package customer

type Address interface {
	AddStreet1(street1 string)
	GetStreet1() (street1 string)

	AddStreet2(street2 string)
	GetStreet2() (street2 string)

	AddCity(city string)
	GetCity() (city string)

	AddState(state string)
	GetState() (state string)

	AddCountry(country string)
	GetCountry() (country string)

	AddPostal(postal string)
	GetPostal() (postal string)

	AddCustomField(key, value string)
	GetCustomField(key string) (value string)

	AddAddressByString(addressString string) (err error)
	Verify() (address Address, isVerified bool, err error)
}
