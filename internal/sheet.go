package internal

import "fmt"

func saveUser(user *User, fileLink string) error {
	fmt.Printf("%v", user)
	return nil
}
