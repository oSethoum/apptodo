// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeTodos holds the string denoting the todos edge name in mutations.
	EdgeTodos = "todos"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TodosTable is the table that holds the todos relation/edge.
	TodosTable = "todos"
	// TodosInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodosInverseTable = "todos"
	// TodosColumn is the table column denoting the todos relation/edge.
	TodosColumn = "user_todos"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "app/backend/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
