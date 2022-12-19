package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "first_name") {
		return errors.New("FisrtName Already Taken")
	}

	if strings.Contains(err, "username") {
		return errors.New("Username Already Taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}

func LoginError (err string) error {
	if strings.Contains(err, "username") {
		return errors.New("Username Already Taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	return errors.New("User not foun!")
}