package main

import (
	"fmt"
	"net/http"
)

func main() {
	//	router = mux.NewRouter()

	fmt.Println("Server running...")
	http.ListenAndServe(":3000", nil)
}
