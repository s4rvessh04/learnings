package structs

import (
	"fmt"
)

type BaseUser struct {
	username string
	password string
}

// Embedded Struct or Promoted Struct
/*
These structs contains the fields of its parent struct + its own fields,
To parent fields can directly be accessed.
*/
type User struct {
	BaseUser
	fullName    string
	email       string
	phoneNumber string
}

func (u User) getBaseUser() BaseUser {
	return BaseUser{
		username: u.username,
		password: u.password,
	}
}

// Nested Struct
/*
These structs contains its own fields + the named parent struct
Inorder to access parent fields we need to access via the parent struct name first.
*/
type Profile struct {
	User     User
	isActive bool
}

func (p Profile) getUsername() string {
	return p.User.username
}

func Test() {
	fmt.Println("Test function inside struct!")

	// Embedded struct invocation
	user1 := User{
		fullName:    "User One",
		email:       "user1@email.com",
		phoneNumber: "12340",
		BaseUser: BaseUser{
			username: "user1",
			password: "password1",
		},
	}

	fmt.Println(User{}.getBaseUser())

	// Embedded struct field access example
	fmt.Printf("Hello %v your email is: %v\n", user1.username, user1.email)

	// Nested struct invocation
	profile1 := Profile{
		User:     user1,
		isActive: true,
	}

	// Nested struct field access example
	fmt.Printf("Hello %v your account active status is: %v\n", profile1.User.username, profile1.isActive)

}
