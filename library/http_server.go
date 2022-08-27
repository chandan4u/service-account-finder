package library

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"service-account-finder/cmd/api/router"
	"syscall"
	"time"

	log "github.com/fatih/color"
)

// Run api server using mux router
func Run(addr string) {
	log.Blue("Listening to port %s ", addr)
	httpServer := &http.Server{
		Addr:    ":" + addr,
		Handler: router.InitializeRoutes(),
	}
	start(httpServer)
}

// Start server on the given port
func start(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Red("Error in Http Server :: ", err)
		}
	}()
	graceFullyShutdown(server)
}

// GracefulShutdown grace fully shutdown for exception case
func graceFullyShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop

	log.Yellow("Shutting the server down.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.Red("Something went wrong in shutdown server :: &s", err)
	} else {
		log.Blue("Server Stop successfully")
	}
}
