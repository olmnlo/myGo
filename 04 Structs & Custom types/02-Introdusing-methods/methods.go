package main

import (
	"example.com/test/user"
)

func main() {
	// Create a new user
	user01 := user.NewUser("Hussam", "Alghamdi", "03/02/2002")

	admin := user.NewAdmin("mm", "kk")
	// Output user details
	user01.Output()
	admin.AdminUser.Output()

	// Clear user data
	user01.ClearUserStruct()

	// Output cleared user details
	user01.Output()
}
