package main

import (
	"github.com/ZachSellman/practice-it-go-rest-api-server/src/backend"
)

func main() {
	a := backend.App{}
	a.Port = ":9003"
	a.Initialize()
	a.Run()
}
