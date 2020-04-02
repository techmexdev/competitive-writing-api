package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/techmexdev/competitive_writing_api/pkg/http/rest"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
	"github.com/techmexdev/competitive_writing_api/pkg/storage/memo"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := memo.New()
	psg := passage.NewService(store)
	usr := user.NewService(store)
	wri := writing.NewService(store)
	sel := selecting.NewService(store)

	logger := log.New()
	logger.Out = os.Stdout

	cfg := rest.Config{Wri: wri, Psg: psg, Usr: usr, Sel: sel}
	handler := rest.New(cfg, logger)

	log.Printf("Starting web server at port = %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
