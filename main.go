/*
 * SEP6-movies backend
 *
 * Backend part of the project delivered for SEP6 course - Movies platform  Authors of project:  Konrad Piotrowski (280053) Aleksander Stefan Bialik (280027)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	"os"
	"studies/SEP6-Backend/swagger"
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
