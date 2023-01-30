package main

import (
	"github.com/ayushjnv1/Gobank/app"
	"github.com/ayushjnv1/Gobank/config"
	"github.com/ayushjnv1/Gobank/server"
)

func main(){
	config.Load()
	app.Init()
	server.StartAPIServer()
}