package models

import (
	"strconv"
	"time"
)

type User struct {
	Id           *string
	Firstname    string
	Lastname     string
	Username     string
	Password     string
	Salt         string
	Address_id   int16
	Created_date time.Time
	Updated_date *time.Time
	Email        string
	Birth_date   time.Time
}

// models should always be immutable
// create implements business an validations
func (u User) Create(
	firstname string,
	lastname string,
	username string,
	password string,
	updated_date *time.Time,
	email string,
	birth_date time.Time,
) User {
	// simple salt,
	// created today
	// addres_id
	create_date := time.Now().UnixMicro()
	epoch := strconv.FormatInt(create_date, 8)
	salt := firstname + lastname + epoch
	return User{nil, firstname, lastname,
		username, password, salt, 0, time.Now(),
		updated_date, email, birth_date}
}
