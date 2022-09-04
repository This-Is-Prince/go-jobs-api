package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/This-Is-Prince/go-jobs-api/db"
	"github.com/This-Is-Prince/go-jobs-api/routes"
)

func main() {
	fmt.Println("Hello World!")

	jobRouter := routes.JobRouter()
	log.Fatal(http.ListenAndServe(":3000", jobRouter))
}
