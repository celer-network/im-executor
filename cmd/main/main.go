/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	_ "embed"
	"fmt"

	"github.com/celer-network/im-executor/cmd"
)

//go:embed logo
var logo string

func main() {
	fmt.Println(logo)
	cmd.Execute()
}
