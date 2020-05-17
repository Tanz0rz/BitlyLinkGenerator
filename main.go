package main

import (
	"encoding/json"
	"fmt"
	"github.com/TannerMoore/BitlyGameLinkGenerator/config"
	"github.com/TannerMoore/BitlyGameLinkGenerator/constants"
	"github.com/TannerMoore/BitlyGameLinkGenerator/requests"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var AuthKey = ""
var Domain = ""

func main() {

	if len(os.Args) != 3 {
		fmt.Println("This tool requires exactly two command line arguments. The first should be your API key and the second should be the domain you want to make the links under")
		os.Exit(-1)
	}

	AuthKey = os.Args[1]
	Domain = os.Args[2]

	if len(AuthKey) == 0 || len(Domain) == 0 {
		fmt.Println("This tool requires exactly two command line arguments. The first should be your API key and the second should be the domain you want to make the links under")
		os.Exit(-1)
	}

	linksBytes, err := ioutil.ReadFile("./config/games.json")
	if err != nil {
		fmt.Printf("Error reading games.json file: %v\n", err)
		os.Exit(1)
	}

	var games config.LinkConfigList
	err = json.Unmarshal(linksBytes, &games)
	if err != nil {
		fmt.Printf("Failed to unmarshal games json: %v\n", err)
	}

	fmt.Printf("Config: %+v\n", games)

	for _, game := range games {
		if len(game.JamPage) > 0 {
			fmt.Printf("Creating jam page links for %v\n", game.Name)
			err := CreateJamPageLink(game.Name, game.JamPage)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			time.Sleep(1 * time.Second)
		}

		if len(game.Github) > 0 {
			fmt.Printf("Creating github page links for %v\n", game.Name)
			err := CreateGameContentLinks(game.Name, game.Github, constants.Github)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			time.Sleep(1 * time.Second)
		}

		if len(game.Ost) > 0 {
			fmt.Printf("Creating ost page links for %v\n", game.Name)
			err := CreateGameContentLinks(game.Name, game.Ost, constants.Ost)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			time.Sleep(1 * time.Second)
		}

		if len(game.ItchIo) > 0 {
			fmt.Printf("Creating itch.io page links for %v\n", game.Name)
			err := CreateGameContentLinks(game.Name, game.ItchIo, constants.ItchIo)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("Link generation completed successfully")
}

func CreateJamPageLink(gameName, jamPage string) error {

	// Create jam page link (only needed once per long url)
	bitlink, err := requests.CreateLink(Domain, jamPage, AuthKey)
	if err != nil {
		return fmt.Errorf("Unable to create jam page link for %v: %v\n", gameName, err)
	}

	// Create the lower case custom backhalf version
	err = requests.SetCustomBackHalf(bitlink, Domain, strings.ToLower(gameName), constants.NoSuffix, AuthKey)
	if err != nil {
		return fmt.Errorf("Unable to create custom bitlink for %v: %v\n", gameName, err)
	}

	// Create the normal case custom backhalf version
	err = requests.SetCustomBackHalf(bitlink, Domain, gameName, constants.NoSuffix, AuthKey)
	if err != nil {
		return fmt.Errorf("Unable to create custom bitlink for %v: %v\n", gameName, err)
	}
	return nil
}

func CreateGameContentLinks(gameName string, linkedContent string, linkSuffixes []string) error {

	// Create content link (only needed once per long url)
	bitlink, err := requests.CreateLink(Domain, linkedContent, AuthKey)
	if err != nil {
		return fmt.Errorf("Unable to create link for %v: %v\n", gameName, err)
	}

	// Create the lower case custom backhalf versions
	for _, linkSuffix := range linkSuffixes {
		err = requests.SetCustomBackHalf(bitlink, Domain, strings.ToLower(gameName), linkSuffix, AuthKey)
		if err != nil {
			return fmt.Errorf("Unable to create custom bitlink for %v: %v\n", gameName, err)
		}
	}

	// Create the normal case custom backhalf versions
	for _, linkSuffix := range linkSuffixes {
		err = requests.SetCustomBackHalf(bitlink, Domain, gameName, linkSuffix, AuthKey)
		if err != nil {
			return fmt.Errorf("Unable to create custom bitlink for %v: %v\n", gameName, err)
		}
	}
	return nil
}
