package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

//GORMBase struct
type GORMBase struct {
	ID        *uuid.UUID  `gorm:"type:uuid;primary_key;json:omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time  `json:"update_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate func
func (base *GORMBase) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate  func
func (base *GORMBase) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
