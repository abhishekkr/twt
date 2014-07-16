package twt

import (
	"github.com/ChimeraCoder/anaconda"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
)

func GetAPI(creds golhashmap.HashMap) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(creds["api-key"])
	anaconda.SetConsumerSecret(creds["api-secret"])
	return anaconda.NewTwitterApi(creds["access-token"], creds["access-token-secret"])
}
