package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Request represents a requests table.
type Request struct {
	ID                string     `gorm:"type:varchar(64); primary_key:true" json:"id"`
	TimeCreated       *time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"time_created"`
	TimeCompleted     *time.Time `json:"time_completed"`
	TimeHoldCreated   *time.Time `json:"-"`
	TimeHoldCompleted *time.Time `json:"-"`
	User              User       `gorm:"ForeignKey:UserID" json:"-"`
	UserID            int64      `gorm:"not null" json:"-"`
}

// RequestCont represents a container for a list of Request records from the requests table.
type RequestCont struct {
	Requests []Request `json:"requests"`
}

// BeforeCreate creates and sets the ID field with a uuid and TimeCreated before creating the record.
func (rr *Request) BeforeCreate(scope *gorm.Scope) error {
	frUUID, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", frUUID.String())
	_ = scope.SetColumn("TimeCreated", time.Now().UTC())
	return nil
}

// Create inserts the Request data into the requests table and fills in the Request ID.
func (rr *Request) Create(db *gorm.DB) error {
	if err := db.Create(&rr).Error; err != nil {
		return err
	}

	return nil
}
