package greetings

import (
	_ "github.com/twitchtv/twirp" // to emulate some package being used
)

//go:generate greetings-gen -pkg x.com/x/greetings

// Hello says hello to Name.
type Hello struct {
	ID   string
	Name string
}

// Goodbye says goodbye to Name.
type Goodbye struct {
	ID   string
	Name string
}
