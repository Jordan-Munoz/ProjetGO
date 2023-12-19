package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

//By Jordan Munoz for ProjectGO 
//Worker = les travailleurs :D
func Worker(target string, dictionary chan string, results chan string) {
	for word := range dictionary {
		// URL en combinant cible et la liste sinon marche pas
		url := fmt.Sprintf("%s/%s", target, word)
		
		// Tester la connexion
		response, err := http.Get(url)
		
		// Verif si la connexion passe ou pas (200 OK - 404 Pas OK)
		if err == nil {
			// Envoie du résultat
			results <- fmt.Sprintf("%s %d", url, response.StatusCode)
		} else {
			// Si échoue = 404
			results <- fmt.Sprintf("%s 404", url)
		}
	}
}

//Scan pour énuméré la liste 
func Scan(target string, dictionaryPath string, workers int, quiet bool) {
	// Ouvrir le fichier du dictionnaire (Merci Stackover)
	file, err := os.Open(dictionaryPath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier du dictionnaire:", err)
		return
	}
	defer file.Close()

	// Créer des chan pour les dict, resultat et la fin
	dictionary := make(chan string, 100)
	results := make(chan string, 100)
	done := make(chan bool)

	var wg sync.WaitGroup

	//Start les workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(target, dictionary, results)
		}()
	}