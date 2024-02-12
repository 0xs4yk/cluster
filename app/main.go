package main

import (
	"fmt"
	"net/http"
)

func handleSmartContract1(w http.ResponseWriter, r *http.Request) {
	// Ajoutez votre logique spécifique pour le Smart Contract 1 ici
	fmt.Fprintf(w, "Traitement du Smart Contract 1")
}

func handleSmartContract2(w http.ResponseWriter, r *http.Request) {
	// Ajoutez votre logique spécifique pour le Smart Contract 2 ici
	fmt.Fprintf(w, "Traitement du Smart Contract 2")
}

func handleSmartContract3(w http.ResponseWriter, r *http.Request) {
	// Ajoutez votre logique spécifique pour le Smart Contract 3 ici
	fmt.Fprintf(w, "Traitement du Smart Contract 3")
}

func setupCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "interface.html")
	})

	http.HandleFunc("/call-smart-contract-1", func(w http.ResponseWriter, r *http.Request) {
		setupCORSHeaders(w, r)
		handleSmartContract1(w, r)
	})

	http.HandleFunc("/call-smart-contract-2", func(w http.ResponseWriter, r *http.Request) {
		setupCORSHeaders(w, r)
		handleSmartContract2(w, r)
	})

	http.HandleFunc("/call-smart-contract-3", func(w http.ResponseWriter, r *http.Request) {
		setupCORSHeaders(w, r)
		handleSmartContract3(w, r)
	})

	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
