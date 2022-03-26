package search

import (
	"fmt"

	UD "github.com/details"
)

func FindUserById() (UD.UserDetail, error) {
	fmt.Println("Enter the user ID")
	var id int
	fmt.Scanln(&id)
	return UD.GetUserDetail(id)
}

func FindMultiUserById() ([]UD.UserDetail, error) {

	fmt.Println("Enter number of  Users to find")
	var num int
	fmt.Scanln(&num)
	ids := make([]int, num)
	for i := 0; i < num; i++ {
		var id int
		fmt.Scanln(&id)
		ids[i] = id
	}
	var listUser []UD.UserDetail
	for _, id := range ids {
		ud, err := UD.GetUserDetail(id)
		if err != nil {
			continue
		}
		listUser = append(listUser, ud)
	}
	if len(listUser) == 0 {
		return []UD.UserDetail{}, fmt.Errorf("no user found")
	}
	return listUser, nil
}

func GetAllUserList() ([]UD.UserDetail, error) {
	lstuser := UD.ListAllUser()
	if len(lstuser) == 0 {
		return lstuser, fmt.Errorf("no user found")
	}
	return lstuser, nil
}
