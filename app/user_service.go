package app

import (
	"log"

	"github.com/agustin-sarasua/rs-model"
)

func CreateUser(u *m.User) (int64, []error) {
	log.Printf("Creating new User: %+v\n", u)
	if errs := validateUser(u); len(errs) > 0 {
		return 0, errs
	}
	Db.Create(u)
	log.Printf("User ID: %+v\n", u.ID)
	return u.ID, nil
}

func LoadUser(uid uint64) *m.User {
	log.Printf("Loading User: %+v\n", uid)
	var u m.User
	Db.First(&u, uid)
	return &u
}

func UpdateUser(u *m.User) (int64, []error) {
	log.Printf("Updating Property: %+v\n", u.ID)
	if errs := validateUser(u); len(errs) > 0 {
		return 0, errs
	}
	Db.Save(u)
	return u.ID, nil
}
