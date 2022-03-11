package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	//fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	a.Listen(port)
	/*if err := a.Listen(port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}*/
}
