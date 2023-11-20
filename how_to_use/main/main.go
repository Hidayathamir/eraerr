package main

import (
	"fmt"

	"github.com/Hidayathamir/eraerr"
	"gorm.io/gorm"
)

func main() {
	err := ServiceDeleteUserIfFraud()
	if err != nil {
		err = eraerr.WrapErr(err, "error service delete user if user fraud")
		fmt.Println(err) // "error service delete user if user fraud:: error delete user:: WHERE conditions required"
		// here, i have full context of the err because i wrap the err.
		// if i do not wrap the err, err will only has value "WHERE conditions required"
		return
	}
}

func ServiceDeleteUserIfFraud() error {
	err := ServiceDetectFraud()
	if err != nil {
		return eraerr.WrapErr(err, "error detect fraud")
	}

	err = RepoDeleteUser()
	if err != nil {
		return eraerr.WrapErr(err, "error delete user")
	}
	return nil
}

func ServiceDetectFraud() error {
	return nil
}

func RepoDeleteUser() error {
	return gorm.ErrMissingWhereClause
}
