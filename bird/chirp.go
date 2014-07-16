package twt

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func PrintUser(user anaconda.User) {
	fmt.Println("Name:", user.Name, "\nIdStr:", user.IdStr, "\nScreenName:", user.ScreenName)
	fmt.Println("\nFollowing:", user.Following, "\nStatus:", user.Status)
	fmt.Println(user.Description)
	fmt.Println("-----------------------------------------------------------------------------")
}
