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
