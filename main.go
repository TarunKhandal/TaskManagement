package main

import (
	"TaskManagement/app/globals"
	"TaskManagement/app/globals/dbconnections"
	"TaskManagement/routes"
	"log"
	"net/http"
	"strings"
)

func main() {

	globals.RootDirPathForAppData = "."
	globals.RootDirPathForAppData = strings.ReplaceAll(globals.RootDirPathForAppData, `\`, `/`)
	if !strings.HasSuffix(globals.RootDirPathForAppData, `/`) {
		globals.RootDirPathForAppData = globals.RootDirPathForAppData + `/`
	}

	globals.InitConfiguration()

	dbconnections.OpenDB()
	defer dbconnections.CloseDB()

	ser := &http.Server{
		Addr:    globals.Configuration.App.Port,
		Handler: routes.Routes(),
	}

	log.Println("Listening and serving HTTP on :" + globals.Configuration.App.Port)
	if err := ser.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
