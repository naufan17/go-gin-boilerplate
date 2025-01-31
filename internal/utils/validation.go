package utils

import "github.com/go-playground/validator/v10"

func ParseValidationError(err validator.ValidationErrors) map[string]string {
	errorMessage := make(map[string]string)

	for _, v := range err {
		switch v.StructField() {
		case "name":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "name is required"
			}
		case "email":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "email is required"
			} else if v.Tag() == "email" {
				errorMessage[v.Field()] = "email is invalid"
			}
		case "password":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "password is required"
			} else if v.Tag() == "min" {
				errorMessage[v.Field()] = "password must be at least 10 characters"
			}
		case "confirm_password":
			if v.Tag() == "required" {
				errorMessage[v.Field()] = "password confirmation is required"
			} else if v.Tag() == "eqfield" {
				errorMessage[v.Field()] = "password confirmation does not match password"
			}
		}
	}

	return errorMessage
}
