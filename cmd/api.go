package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edmiltonVinicius/register-steps/domain"
)

func main() {

	domain.LoadEnv()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello word!!! \n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}