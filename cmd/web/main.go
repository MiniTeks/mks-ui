// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/MiniTeks/mks-ui/pkg/db"
	"github.com/go-redis/redis/v8"
)

var (
	dbAddr = flag.String("dbAddr", "127.0.0.1:6379", "The address of the redis server")
	dbPass = flag.String("password", "12345", "The password of the Kubernetes API server")
	addr   = flag.String("addr", ":5000", "Port Value")
)

var rClient *redis.Client

func main() {

	flag.Parse()

	// redis-db client
	cred := db.RClient{
		Addr: *dbAddr,
		Pass: *dbPass,
		Db:   0,
	}

	rClient = db.GetRedisClient(&cred)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomePage)

	// fileServer := http.FileServer(http.Dir("./ui/static"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on localhost%s\n", *addr)
	er := http.ListenAndServe(*addr, mux)
	log.Fatal(er)
}
