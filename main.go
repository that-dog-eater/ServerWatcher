package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"mvp/backend/api"
	servermanager "mvp/backend/server_manager"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

var outputDir = "C:\\Users\\maint\\Desktop\\Dev\\Projects\\Sandbox\\startup\\mvp\\backend\\output"
var servers_file_path = "C:\\Users\\maint\\Desktop\\Dev\\Projects\\Sandbox\\startup\\mvp\\backend\\output\\servers.json"

func main() {

	//Backend
	servers, err := servermanager.LoadServers(servers_file_path)
	if err != nil {
		fmt.Println("Loading servers failed: ", err)
	}

	for _, srv := range servers {
		fmt.Println("Snapshotting started for server: ", srv.Name)
		servermanager.PerServerTask(srv, outputDir)
	}

	// Frontend

	apiInstance := api.NewAPI(outputDir, servers_file_path)

	app := application.New(application.Options{
		Name:        "mvp",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(apiInstance),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Window 1",
		Width:  1000,
		Height: 800,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
