package service

import (
	"fmt"
	"testing"
)

func TestUserJson(t *testing.T) {
	user := GetUser(1, 3)
	fmt.Println(user)
}
