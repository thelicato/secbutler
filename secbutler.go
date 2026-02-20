package main

import (
	"github.com/thelicato/secbutler/cmd"
	"github.com/thelicato/secbutler/pkg/utils"
)

func main() {
	utils.Banner(utils.Version)
	cmd.Execute()
}
