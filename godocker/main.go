package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	appVersion = "n/a"
	appCommit  = "n/a"
	appBuilt   = "n/a"
)

func main() {
	version := flag.Bool("v", false, "prints current app version")
	flag.Parse()
	if *version {
		fmt.Printf("Version : %v \nCommit : %v\nBuilt: %v\n", appVersion, appCommit, appBuilt)
		os.Exit(0)
	}

	router := mux.NewRouter()
	router.PathPrefix("/app").Handler(http.StripPrefix("/app", http.FileServer(http.Dir("web/app/dist"))))
	router.PathPrefix("/").HandlerFunc(indexHandler)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Docker :D\nThis is awesome :)")
}
