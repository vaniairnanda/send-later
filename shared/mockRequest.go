package shared

import "errors"

func MockBusinessServiceGetUserLogin(login string) error {
	if login == "WRONG_LOGIN" {
		return errors.New("you do not have access to complete this action")
	}
	return nil
}
