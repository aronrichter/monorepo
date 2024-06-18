package main

import (
	"fmt"

	"github.com/aronrcihter/monorepo/applications/bankdomicile/gateway/usermetrics"
	"github.com/aronrcihter/monorepo/libs/client"
)

func main() {
	c := usermetrics.NewClient(client.Client{})

	fmt.Println(c)
}
