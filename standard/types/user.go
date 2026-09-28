package types

import (
	"fmt"
	"uuid"
)

// In Go, types are defined as struct, but in many other languages, it is called class.
// The fields are public by default if the FIRST LETTER is UPPERCASE
// and private if the FIRST LETTER is LOWERCASE
type User struct {
	Id       uuid.UUID // Public field
	Name     string    // Public field
	Email    string    // Public field
	password string    // Private field
}

// Creates a new [User].
//
// It receives a pointer to an IdGenerator to generate the ID of the new user
// and increments it for the next user.
func NewUser(name string, email string, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		password: password,
	}
}

// By using the * syntax, we can reference a field from the struct without passing the struct as an argument
// This is a more efficient way to reference a field from the struct
// This is called a "pointer receiver"
func (user *User) GetName() string {
	return user.Name
}

func (user *User) GetID() uuid.UUID {
	return user.Id
}

// Change the name of a user
func (user *User) ChangeName(name string) {
	user.Name = name
}

// Change the password of a user
func (user *User) ChangePassword(password string) {
	user.password = password
}

// Change the email of a user
func (user *User) ChangeEmail(email string) {
	user.Email = email
}

// String method to print the user
func (user *User) String() string {
	return fmt.Sprintf("{\n\tId: \t%s\n\tName: \t%s\n\tEmail: \t%s\n}", user.Id, user.Name, user.Email)
}
