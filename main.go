package main

import (
	"fmt"
	"subteez/config"
	"subteez/subteez"
)

func main() {
	appConfig := config.NewConfigFile("subteez.config")
	if appConfig.Load() != nil {
		appConfig.SetServer("https://subteez1.herokuapp.com")
		appConfig.SetLanguageFilters(subteez.Languages)
		if appConfig.Save() != nil {
			fmt.Println("Note: Error in saving default config")
		}
	}
}
