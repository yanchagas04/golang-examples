package types

// Represents a simple address
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// NewAddress creates a new Address
func NewAddress(street, city, state, zip string) Address {
	return Address{
		Street: street,
		City:   city,
		State:  state,
		Zip:    zip,
	}
}

// In go, struct composition is called embedding.
// This is a way to include the fields of one struct into another struct.
type Person struct {
	Name    string
	Email   string
	Address Address
}

// NewPerson creates a new Person by creating a new Address and embedding it in the Person struct
func NewPerson(name, email, street, city, state, zip string) Person {
	return Person{
		Name:    name,
		Email:   email,
		Address: NewAddress(street, city, state, zip),
	}
}
