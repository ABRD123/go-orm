package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// User represents an users table.
type User struct {
	ID          int64      `gorm:"primary_key" json:"-"`
	TimeCreated *time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"-"`
	TimeUpdated *time.Time `json:"-"`
	Name        string     `gorm:"type:varchar(128); unique; not null" json:"name"`
	Active      bool       `gorm:"type:bool; default:false" json:"-"`
}

// BeforeCreate sets the TimeCreated and TimeUpdated before creating the record.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	imageUUID, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", imageUUID.String())
	return nil
}

// BeforeSave sets the TimeUpdated before saving the record.
func (user *User) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("time_updated", time.Now().UTC())
	return nil
}

// BeforeUpdate sets the TimeUpdated before updating the record.
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("time_updated", time.Now().UTC())
	return nil
}

// Create inserts the user data into the users table and fills in the Image ID.
func (user *User) Create(db *gorm.DB) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetActiveUserName gets a User record by Active and UserName.
//	param: active is the User Active
//	param: userName is the User UserName
func (user *User) GetActiveUserName(active bool, userName string, db *gorm.DB) error {
	query := "active = ? and name = ?"
	if err := db.Where(query, active, userName).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// GetUserName gets a User record by UserName.
//	param: userName is the User UserName
func (user *User) GetUserName(userName string, db *gorm.DB) error {
	query := "name = ?"
	if err := db.Where(query, userName).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateActive updates the User Active value in the DB.
//	param: active is the User Active
func (user *User) UpdateActive(active bool, db *gorm.DB) error {
	if err := db.Model(u).Update("active", active).Error; err != nil {
		return err
	}
	return nil
}
