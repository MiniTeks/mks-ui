package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/MiniTeks/mks-ui/pkg/db"
)

var (
	dbAddr = flag.String("dbAddr", "127.0.0.1:6379", "The address of the redis server")
	dbPass = flag.String("password", "12345", "The password of the Kubernetes API server")
	addr   = flag.String("addr", ":4000", "Port Value")
)

func main() {

	flag.Parse()

	// redis-db client
	cred := db.RClient{
		Addr: *dbAddr,
		Pass: *dbPass,
		Db:   0,
	}
	rClient := db.GetRedisClient(&cred)
	dapp, err := db.GetValues(rClient)
	// fmt.Print(dapp)
	app := (*wrap)(dapp)
	fmt.Print(app)
	if err != nil {
		log.Fatalf("Couldn't get the values from the source")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomePage)

	// fileServer := http.FileServer(http.Dir("./ui/static"))

	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s\n", *addr)
	er := http.ListenAndServe(*addr, mux)
	log.Fatal(er)
}
