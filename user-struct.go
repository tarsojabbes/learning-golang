package main

import (
	"errors"
	"fmt"
)

type User struct {
	name  string
	email string
	age   int
}

func main() {

	first_user := User{"John Doe", "johndoe@email.com", 18}
	first_user_created, error_first_user := create_user(first_user)

	if error_first_user != nil {
		fmt.Println(error_first_user.Error())
	} else {
		fmt.Println(first_user_created)
	}

	second_user := User{"", "johndoe@email.com", 11}
	second_user_created, error_second_user := create_user(second_user)

	if error_second_user != nil {
		fmt.Println(error_second_user.Error())
	} else {
		fmt.Println(second_user_created)
	}

}

func create_user(user User) (User, error) {
	if user.age > 13 && user.email != "" && user.name != "" {
		return user, nil
	} else {
		return user, errors.New("An error occured to create this user")
	}
}
