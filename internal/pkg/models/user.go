package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User represents a users table.
type User struct {
	ID          int64      `gorm:"primary -key" json:"-"`
	TimeCreated *time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"-"`
	TimeUpdated *time.Time `json:"-"`
	Name        string     `gorm:"type:varchar(128); unique; not null" json:"name"`
	Active      bool       `gorm:"type:bool; default:false" json:"-"`
}

// BeforeCreate sets the TimeCreated and TimeUpdated before creating the record.
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("time_created", time.Now().UTC())
	_ = scope.SetColumn("time_updated", time.Now().UTC())
	return nil
}

// BeforeSave sets the TimeUpdated before saving the record.
func (u *User) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("time_updated", time.Now().UTC())
	return nil
}

// BeforeUpdate sets the TimeUpdated before updating the record.
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("time_updated", time.Now().UTC())
	return nil
}

// Create inserts the User data into the users table and fills in the User ID.
func (u *User) Create(db *gorm.DB) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

// GetActiveUserName gets a User record by Active and UserName.
//	param: active is the User Active
//	param: userName is the User UserName
func (u *User) GetActiveUserName(active bool, userName string, db *gorm.DB) error {
	query := "active = ? and name = ?"
	if err := db.Where(query, active, userName).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// GetUserName gets a User record by UserName.
//	param: userName is the User UserName
func (u *User) GetUserName(userName string, db *gorm.DB) error {
	query := "name = ?"
	if err := db.Where(query, userName).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateActive updates the User Active value in the DB.
//	param: active is the User Active
func (u *User) UpdateActive(active bool, db *gorm.DB) error {
	if err := db.Model(u).Update("active", active).Error; err != nil {
		return err
	}
	return nil
}
