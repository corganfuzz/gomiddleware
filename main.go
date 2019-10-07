package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase...")

		//Pass control back to handler
		handler.ServeHTTP(w, r)

		fmt.Println("Executing middleware after request phase...")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Biz logic goes here

	fmt.Println("Executing main handler...")
	w.Write([]byte("OK"))
}

func main() {
	//HandlerFunc return a HTTP Handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server is running in port 8000...")
}
