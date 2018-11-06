package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gobuffalo/packr"
)

var user = syscall.NewLazyDLL("user32.dll")

func main() {
	// Store files for the interface (packed into binary) in the interface folder
	box := packr.NewBox("./interface")

	// Serve static files from interface folder
	http.Handle("/", http.FileServer(box))
	// Handle killswitch endpoint
	http.HandleFunc("/kill", buttonHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting interface server: %v\n", err)
		os.Exit(1)
	}
}

func lockScreen() {
	proc := user.NewProc("LockWorkStation")
	_, _, _ = proc.Call()
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	lockScreen()
	w.Header().Set("Location", "/?success")
	w.WriteHeader(302)
	fmt.Fprintf(w, "Redirecting")
}
