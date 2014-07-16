package main

import (
	"flag"
	"fmt"

	golconfig "github.com/abhishekkr/gol/golconfig"
	golhashmap "github.com/abhishekkr/gol/golhashmap"

	twt "github.com/abhishekkr/twt/bird"
)

var (
	config = flag.String("config", "", "config json file for API Keys")
)

func panic_for_config(creds golhashmap.HashMap) {
	if creds["api-key"] == "" || creds["api-secret"] == "" ||
		creds["access-token"] == "" || creds["access-token-secret"] == "" {
		panic("Config doesn't have all credentials required.")
	}
}

func main() {
	var creds golhashmap.HashMap
	flag.Parse()

	creds = make(golhashmap.HashMap)
	json := golconfig.GetConfig("json")
	json.ConfigFromFile(*config, &creds)

	panic_for_config(creds)

	fmt.Println("starting stalking")
	twt.StalkFollowers(creds)
}
