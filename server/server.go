package server

import (
	"net/http"

	"net"

	"github.com/sirupsen/logrus"
	_ "gopkg.in/workanator/vuego.v1/resource"
	"gopkg.in/workanator/vuego.v1/ui"
)

type Server struct {
	ListenIP   net.IP
	ListenPort uint16
	Router     http.Handler
	log        *logrus.Entry
	screens    []ui.Screener
}

// Start prepares the server instance and starts listen for incoming requests. If some fields in the struct omitted
// they are initialized with default values. In most cases it's enough to provide valid ListenIP and ListenPort
// to start the server. When server started it blocks further execution of the current goroutine.
func (server Server) Start(screen ui.Screener) error {
	// Prepare the instance for start.
	if err := server.prepare(); err != nil {
		return err
	}

	// Start the server
	server.log.Info("Starting server")
	err := http.ListenAndServe(":8008", server.Router)

	// Ignore server closed error.
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Prepares the server for start.
func (server *Server) prepare() error {
	// Use teh default router if no router is provided.
	if server.Router == nil {
		server.Router = DefaultRouter()
	}

	// Use the standard logger if no logger is provided.
	if server.log == nil {
		server.log = logrus.NewEntry(logrus.StandardLogger())
	}

	// Initialize screen stack.
	if server.screens == nil {
		server.screens = make([]ui.Screener, 0)
	}

	return nil
}
