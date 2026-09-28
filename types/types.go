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

// To create Methods for a "Class" in Go that doesn't modify the original struct, just use the default function syntax and pass the struct as an argument to reference a field from the "object".
//
// * This is not the efficient way to reference a field from the struct
// * This is called a "value receiver"
func GetName(user User) string {
	return user.Name
}

// By using the * (pointer) syntax, we can reference a field from the struct without passing the struct as an argument
// This is a more efficient way to reference a field from the struct
// This is called a "pointer receiver"
func GetID(user *User) int {
	return user.ID
}

// If we want to modify a value inside a struct, we need to use the * (pointer) syntax
func ChangeEmail(user *User, email string) {
	user.Email = email
}

// The same logic applies to private fields
func ChangePassword(user *User, password string) {
	user.password = password
}
