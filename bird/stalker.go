package twt

import (
	"fmt"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
)

func StalkFollowers(creds golhashmap.HashMap) {
	api := GetAPI(creds)
	pages := api.GetFollowersListAll(nil)
	for page := range pages {
		if page.Error != nil {
			fmt.Println("Aaaaah! Follower stalker failed.\n", page.Error)
			continue
		}
		for _, follower := range page.Followers {
			PrintUser(follower)
		}
	}
}
