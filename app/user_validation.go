package app

import (
	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func validateUser(u *m.User) []error {
	var errs []error
	errs = c.ValidateCondition(func() bool { return u.Name != "" }, "Name can not be empty", errs)
	errs = c.ValidateCondition(func() bool { return u.Email != "" }, "Email can not be empty", errs)
	errs = c.ValidateCondition(func() bool { return u.FirebaseID != "" }, "FirebaseID can not be empty", errs)
	return errs
}
