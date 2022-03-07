package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/pedrocmart/canvas/cmd/core/config"
	"github.com/pedrocmart/canvas/internal/core"

	httpHandlers "github.com/pedrocmart/canvas/internal/core/http"
	middleware "github.com/pedrocmart/canvas/internal/core/http/middleware"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
)

const errorChan int = 10

func main() {

	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbStore, err := config.NewStore(c)
	if err != nil {
		log.Fatal(err)
	}

	service := core.New(c, dbStore)

	h := httpHandlers.NewHandler(c, service)

	router := httptrace.New()

	router.POST("/canvas/add", h.CanvasCreate)
	router.GET("/canvas", h.CanvasGet)

	var (
		httpAddr = "0.0.0.0:8080"
	)

	if port := os.Getenv("CANVAS_PORT"); port != "" {
		httpAddr = fmt.Sprintf("0.0.0.0:%s", port)
	}

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = middleware.ReqID(middleware.Logger(c, router))

	errChan := make(chan error, errorChan)

	go func() {
		c.Log.Debugf("HTTP service listening on %s", httpAddr)
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				c.Log.Fatal(err)
			}
		case s := <-signalChan:
			c.Log.Debugf(fmt.Sprintf("Captured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
