package util

import "net/http"

func NewContactForm(onSubmit http.HandlerFunc) *SimpleForm {
	f := &SimpleForm{
		TitleText: "Contact",
		Fields: []Field{
			{
				Name: NewName("First name"),
				Type: "first-name",
			},
			{
				Name: NewName("Last name"),
				Type: "last-name",
			},
			{
				Name: NewName("Email"),
				Type: "email",
			},
			{
				Name: NewName("Message"),
				Type: "text",
			},
		},
		SubmitText: "Send",
		HandlePOST: onSubmit,
	}
	return f
}
