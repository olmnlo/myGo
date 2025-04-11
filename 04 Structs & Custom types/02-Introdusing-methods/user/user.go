package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName, LastName, BirthDate string
	CreatedAt                      time.Time
}
type Admin struct {
	Email     string
	Phone     string
	AdminUser User
}

// NewUser creates a new User instance with the provided data and current timestamp
func NewUser(userFirstName, userLastName, userBirthDate string) User {
	return User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreatedAt: time.Now(),
	}
}

func NewAdmin(email, phone string) Admin {
	return Admin{
		Email: email,
		Phone: phone,
		AdminUser: User{
			FirstName: "Admin",
			LastName:  "___",
			BirthDate: "___",
			CreatedAt: time.Now(),
		},
	}
}

// Output prints the user's information (FirstName, LastName, BirthDate, and CreatedAt)
func (u *User) Output() {
	fmt.Printf("Name: %s %s\nBirth Date: %s\nCreated At: %s\n",
		u.FirstName, u.LastName, u.BirthDate, u.CreatedAt.Format("2006-01-02 15:04:05"))
}

// ClearUserStruct clears the user's information
func (u *User) ClearUserStruct() {
	u.FirstName = ""
	u.LastName = ""
	u.BirthDate = ""
	u.CreatedAt = time.Now() // Optionally, reset the created time
}
