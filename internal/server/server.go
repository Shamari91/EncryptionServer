package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Shamari91/EncryptionServer/internal/handler"
	"github.com/julienschmidt/httprouter"
)

const (
	bindAddr string = ":8080"
	encrypt  string = "/encrypt"
	decrypt  string = "/decrypt"
)

var (
	httpServer *http.Server
)

func Init() {
	router := httprouter.New()

	router.POST(encrypt, handler.Encrypt)
	router.POST(decrypt, handler.Decrypt)
	httpServer = &http.Server{
		Addr:         bindAddr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}
}

func Run() {
	log.Printf("starting server on %s", bindAddr)
	log.Fatalf("error :: %v", httpServer.ListenAndServe())
}
