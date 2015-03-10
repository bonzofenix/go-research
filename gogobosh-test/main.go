package main

import (
	"fmt"

	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/local"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/cloudfoundry-community/gogobosh/utils"
)

func main() {
	utils.Logger = utils.NewLogger()

	configPath, err := local.DefaultBoshConfigPath()

	if err != nil {
		return
	}

	config, err := local.LoadBoshConfig(configPath)
	if err != nil {
		return
	}

	target, username, password, err := config.CurrentBoshTarget()
	if err != nil {
		return
	}

	director := gogobosh.NewDirector(target, username, password)

	repo := api.NewBoshDirectorRepository(&director, net.NewDirectorGateway())

	info, apiResponse := repo.GetInfo()
	if apiResponse.IsNotSuccessful() {
		fmt.Println("Could not fetch BOSH info")
		return
	}

	fmt.Println("Director")
	fmt.Printf("  Name       %s\n", info.Name)
	fmt.Printf("  URL        %s\n", info.URL)
	fmt.Printf("  Version    %s\n", info.Version)
	fmt.Printf("  User       %s\n", info.User)
	fmt.Printf("  UUID       %s\n", info.UUID)
	fmt.Printf("  CPI        %s\n", info.CPI)
}
