package main

import (
	"fmt"

	"github.com/rabih/go-blog/accountservice/dbclient"
	"github.com/rabih/go-blog/accountservice/service"
)

var (
	appName  = "accountservice"
	DBCleint dbclient.IBoltClient
)

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
