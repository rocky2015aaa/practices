package main

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	users := []string{"user5", "user6", "user7"}
	for index := range users {
		token := (generateToken(users[index]).Value)
		tokenlen := len(token)
		if tokenlen == 0 {
			t.Error("Expected token but got nothing")
		}
	}
}
