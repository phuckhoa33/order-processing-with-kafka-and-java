package base

import "time"

// AuditModelAbstract is a struct meant to be embedded in other model structs
// to provide them with created and updated timestamp fields. These fields
// are automatically populated and updated to reflect the model's lifecycle.
type AuditModelBase struct {
	*ModelBase
	// CreatedAt stores the timestamp of when the model instance was created.
	// It's automatically set upon insertion into the database.
	CreatedAt time.Time `json:"createdAt" gorm:"type:date;"`
	// UpdatedAt stores the timestamp of the last update made to the model instance.
	// It's automatically updated on any change to the model in the database.
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:date;"`
}
