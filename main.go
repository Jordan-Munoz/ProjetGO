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

	//Lire le dict et donner aux workers
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dictionary <- scanner.Text()
		}
		close(dictionary)
	}()

	// Afficher les résultats pendant le scan
	go func() {
		for result := range results {
			if !quiet {
				fmt.Println(result)
			}
		}
		done <- true // Fin du scan
	}()

	fmt.Println("Starting scan...")

	// Attendre la fin du scan des workers
	wg.Wait()

	// Fermer le canal des résultats
	close(results)

	// Attendre la fin du traitement des résultats (Merci mec de stack)
	<-done

	// Afficher le temps écoulé
	fmt.Println("Scan terminé en", time.Since(start))
}

var start time.Time

func main() {
    // Définir les flag 
    dictionaryPath := flag.String("d", "", "Chemin vers le fichier du dictionnaire")
    quiet := flag.Bool("q", false, "Mode silencieux, affiche uniquement le HTTP 200")
    target := flag.String("t", "", "Cible à énumérer")
    workers := flag.Int("w", 1, "Nombre de travailleurs à exécuter")

    // Analyser les flags
    flag.Parse()

    // Vérifier si les flags obligatoires sont définis
    if *dictionaryPath == "" || *target == "" {
        fmt.Println("Utilisation de mygb:")
        flag.PrintDefaults()
        return
    }

    // Initialiser la variable start
    start = time.Now()

    // Afficher les détails de la configuration
    fmt.Println("Démarrage de MyGB")
    fmt.Println("---")
    fmt.Println("Cible:", *target)
    fmt.Println("Liste:", *dictionaryPath)
    fmt.Println("Travailleurs:", *workers)
    fmt.Println("---")

    // Démarrer le scan
    Scan(*target, *dictionaryPath, *workers, *quiet)
}
