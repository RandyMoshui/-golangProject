package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFindUsersByUsername(t *testing.T) {
	user, err := FindUsersByUsername("luke")
	fmt.Println("user: ", user,"err: ",err)
}

func TestAllUsers(t *testing.T) {
	Users, err := AllUsers()
	fmt.Println("user: ", Users, "err:", err)
}

func TestFindRoleById(t *testing.T) {
	role := FindRoleById(1)
	b, _ := json.MarshalIndent(role, "", "     ")
	fmt.Println(string(b))
}


