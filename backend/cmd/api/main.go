package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/willholmeswastaken/wubb/internal/handlers"
	"net/http"
)

func main() {
	log.SetReportCaller(true)

	mux := http.NewServeMux()

	handlers.Handler(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started on port 8080")
}
