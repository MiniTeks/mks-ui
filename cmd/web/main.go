package main

import (
	"flag"
	"log"
	"net/http"
)

type mksresource struct {
	Created   int
	Active    int
	Failed    int
	Completed int
	Deleted   int
}

type application struct {
	Mkstask, Mkspipelinerun, Mkstaskrun mksresource
}

func main() {

	addr := flag.String("addr", ":4000", "Port Value")
	flag.Parse()

	app := &application{
		Mkstask: mksresource{
			Created:   12,
			Active:    23,
			Failed:    34,
			Completed: 45,
			Deleted:   56,
		},
		Mkspipelinerun: mksresource{
			Created:   12,
			Active:    23,
			Failed:    34,
			Completed: 45,
			Deleted:   56,
		},
		Mkstaskrun: mksresource{
			Created:   12,
			Active:    23,
			Failed:    34,
			Completed: 45,
			Deleted:   56,
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homePage)

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
