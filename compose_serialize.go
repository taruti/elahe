package main

import (
	"net/mail"
)

func isValidAddress(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func isValidAddressList(s string) bool {
	_, err := mail.ParseAddressList(s)
	return err == nil
}
