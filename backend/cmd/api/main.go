package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/willholmeswastaken/wubb/internal/handlers"
)

func main() {
	log.SetReportCaller(true)

	r := chi.NewRouter()

	handlers.Handler(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started on port 8080")
}
