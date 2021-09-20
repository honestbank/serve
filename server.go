package serve

import (
	"log"
	"net/http"
)

func (a app) Start() {
	log.Println("starting at localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", a.router))
}
