package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/workanator/vuego.v1/browser"
	"gopkg.in/workanator/vuego.v1/server"
	"gopkg.in/workanator/vuego.v1/ui"
)

func main() {
	// Configure the logger
	logrus.SetLevel(logrus.DebugLevel)

	// Start the browser
	if err := browser.Lauch(
		"http://127.0.0.1:8008/app.html",
		&browser.Options{
			NewInstance: false,
			Window: &ui.WindowOptions{
				Size: &ui.BoxSize{
					Width: 800,
				},
			},
		},
		browser.Firefox(),
	); err != nil {
		logrus.WithError(err).Fatal("Failed to start application")
	}

	// Start the server
	logrus.Fatal(server.DefaultServer().Start(nil))
}
