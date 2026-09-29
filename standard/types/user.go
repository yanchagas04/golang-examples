package types

import (
	"errors"
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

// Set the name of a user
func (user *User) SetName(name string) {
	user.Name = name
}

// Set the password of a user
func (user *User) SetPassword(password string) {
	user.password = password
}

// Set the email of a user
func (user *User) SetEmail(email string) {
	user.Email = email
}

// String method to print the user
func (user *User) String() string {
	return fmt.Sprintf("{\n\tId: \t%s\n\tName: \t%s\n\tEmail: \t%s\n}", user.Id, user.Name, user.Email)
}

// Greet a user
//
// It receives the name of the other person to greet and returns the greeting message
// and an error if the other person's name is empty
func (user *User) Greet(other_person string) (string, error) {
	if other_person == "" {
		return "", errors.New("The other person's name cannot be empty!")
	}
	return fmt.Sprintf("Hello, %s, from %s!", other_person, user.Name), nil
}
