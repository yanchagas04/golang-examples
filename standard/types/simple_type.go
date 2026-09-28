package types

// In Go, types are defined as struct, but in many other languages, it is called class.
// The fields are public by default if the FIRST LETTER is UPPERCASE
// and private if the FIRST LETTER is LOWERCASE
type User struct {
	ID       int
	Name     string // Public field
	Email    string // Public field
	password string // Private field
}

// GetName returns the user's name.
//
// It receives a [User] by value, which creates a copy and ensures the original
// struct cannot be modified.
//
// Notes:
//  - Passing by value is less memory-efficient for large structs than passing a pointer.
//  - When defined as a method `func (u User) GetName() string`, this is known as a value receiver.
func GetName(user User) string {
	return user.Name
}

// By using the * syntax, we can reference a field from the struct without passing the struct as an argument
// This is a more efficient way to reference a field from the struct
// This is called a "pointer receiver"
func GetID(user *User) int {
	return user.ID
}

// Change the name of a user
func ChangeName(user *User, name string) {
	user.Name = name
}

// Change the password of a user
func ChangePassword(user *User, password string) {
	user.password = password
}

// Change the email of a user
func ChangeEmail(user *User, email string) {
	user.Email = email
}

// Ccreates a new [User].
//
// In many other languages, this is called a constructor function.
// But, in Go, we don't have constructor functions, instead, we use the struct literal syntax to create a new struct.
//
// We do this to provide a way to create a new struct in a more organized way.
func NewUser(id int, name string, email string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		password: password,
	}
}