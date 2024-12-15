// Package base defines the foundational structures for model abstraction
// within the application. It includes interfaces and base structs that
// other model structs throughout the application can embed to inherit
// common fields and functionality.
package base

import "github.com/google/uuid"

// IModel defines an interface for audit models. It's used to ensure
// that structs embedding the AuditModelAbstract struct adhere to a basic
// audit model contract, although it's not strictly enforced through methods
type IModel interface{}

// ModelAbstract provides a base struct with a universally unique identifier (UUID)
// as the primary key. It's intended to be embedded in other model structs
// to provide them with an ID field and corresponding database annotations.
type ModelBase struct {
	// ID is the unique identifier for the model, automatically generated as a UUID.
	// It serves as the primary key in the database.
	IModel
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
}
