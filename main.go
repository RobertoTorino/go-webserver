package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func getwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		os.Exit(1)
	}
	return dir
}

func executeScript(w http.ResponseWriter, r *http.Request) {
	// Enable CORS (Cross-Origin Resource Sharing)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userChoice := r.FormValue("choice")
	command := exec.Command("sh", "-c", fmt.Sprintf("echo %s | ./profiles.sh", userChoice))

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Script output: %s\n", output)
	w.WriteHeader(http.StatusOK)
}

func main() {
	currentWorkingDir := getwd()
	fmt.Println("Current working directory:", currentWorkingDir)

	// Dynamically extract the port from the URL
	port := 3000 // or use the value you extracted from the URL
	addr := fmt.Sprintf(":%d", port)

	http.HandleFunc("/executeScript", executeScript)

	server := &http.Server{
		Addr: addr,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Print the dynamically assigned port
	fmt.Printf("Server listening on port %s\n", server.Addr)

	// Wait for an interrupt signal to gracefully shut down the server
	select {}
}
