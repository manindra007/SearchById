package main

import (
	"fmt"

	search "github.com/search"
	uops "github.com/userops"
)

func main() {

	for {
		fmt.Println("select the  operation")
		fmt.Println("1. Add user")
		fmt.Println("2. Update User")
		fmt.Println("3. Search User")
		fmt.Println("4. Search Multiple Users")
		fmt.Println("5. Delete User")
		fmt.Println("6. Delete All  Users")
		fmt.Println("7. Show All Users")
		fmt.Println("e. exit")
		var val int
		fmt.Scanln(&val)
		switch val {
		case 1:
			uops.CreateUser()
		case 2:
			if err := uops.UpdateUserDetail(); err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("User Detail Updated Successfully")
		case 3:
			sr, err := search.FindUserById()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(sr)
		case 4:
			lst, err := search.FindMultiUserById()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(lst)
		case 5:
			err := uops.DeleteUser()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("User Deleted Successfully")
		case 6:
			err := uops.DeleteAllUsers()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("All User Deleted Successfully")
		case 7:
			users, err := search.GetAllUserList()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(users)

		default:
			fmt.Println("exitting")
			return

		}

	}

	// uops.CreateUser(1, "manindra", "gorakhpur", "8299661294", 160, false)
	// uops.CreateUser(2, "mani", "gorakhpur", "8299661294", 160, false)
	// uops.CreateUser(3, "mns", "gorakhpur", "8299661294", 160, false)
	// uops.CreateUser(4, "Singh", "gorakhpur", "8299661294", 160, false)

}
