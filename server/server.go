package server

import (
	"fmt"
	"strconv"

	"github.com/ayushjnv1/Gobank/config"
	"github.com/urfave/negroni"
)

func StartAPIServer(){
	port := config.AppPort()
	server := negroni.Classic()
    dependency := initDependency()
	router := initRouter(dependency) 
	server.UseHandler(router)
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))	
	server.Run(addr)
}