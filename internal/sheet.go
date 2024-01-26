package internal

import "fmt"

func saveUser(user *User) error {
	fmt.Printf("%v", user)
	return nil
}
