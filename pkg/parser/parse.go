package parser

import (
	"flag"
)

var (
	InputFlags = &Inputs{}
)

type Inputs struct {
	Email string
	ID    string
	Token string
}

func Parse() {
	// Define flags
	email := flag.String("email", "", "Your email")
	id := flag.String("userID", "", "Your PD ID")
	token := flag.String("apiToken", "", "Your PD API Token")

	// Parse flags
	flag.Parse()

	InputFlags.Email = *email
	InputFlags.ID = *id
	InputFlags.Token = *token
}
