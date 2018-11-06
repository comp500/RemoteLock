package main

/*
#undef _WIN32_WINNT
#define _WIN32_WINNT 0x0601

#include <stdio.h>
#include <stdlib.h>
#include <windows.h>
#include <strsafe.h>

void Lock() {
	LockWorkStation();
}
*/
import "C"
import (
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	// Store files for the interface (packed into binary) in the interface folder
	box := packr.NewBox("./interface")

	// Serve static files from interface folder
	http.Handle("/", http.FileServer(box))
	// Handle websocket endpoints
	http.HandleFunc("/kill", buttonHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting interface server: %v\n", err)
		os.Exit(1)
	}

}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	C.Lock();
	w.Header().Set("Location", "/")
	w.WriteHeader(302)
	fmt.Fprintf(w, "Redirecting")
}
