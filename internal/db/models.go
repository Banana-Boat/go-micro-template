// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UsersGender string

const (
	UsersGenderFemale UsersGender = "female"
	UsersGenderMale   UsersGender = "male"
)

func (e *UsersGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersGender(s)
	case string:
		*e = UsersGender(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersGender: %T", src)
	}
	return nil
}

type NullUsersGender struct {
	UsersGender UsersGender
	Valid       bool // Valid is true if UsersGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersGender) Scan(value interface{}) error {
	if value == nil {
		ns.UsersGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersGender), nil
}

type User struct {
	ID        int32       `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	Gender    UsersGender `json:"gender"`
	Age       int32       `json:"age"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
