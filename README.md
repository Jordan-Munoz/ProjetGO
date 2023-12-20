# GoBousester - Outil d'énumération de répertoires web en Go

GoBousester est un programme en Go conçu pour effectuer une énumération de répertoires sur un serveur web cible. Il fonctionne en testant différentes URL avec des mots du dictionnaire et en analysant les codes de réponse HTTP pour déterminer si les fichiers ou répertoires existent.

## Utilisation
go run main.go -t <cible> -d <chemin_dictionnaire> -w <nombre_travailleurs> [-q]

## Options

-t: Spécifie la cible à énumérer (ex: http://localhost:8080).
-d: Chemin vers le fichier du dictionnaire.
-w: Nombre de travailleurs à exécuter en parallèle (par défaut: 1).
-q: Mode silencieux, n'affiche que les résultats HTTP 200.

## Exemple

go run main.go -t http://localhost:8080 -d /usr/share/dicts/sec.list -w 5

## Fonctionnalités
Utilise des goroutines pour une exécution concurrente et rapide.
Prend en charge une liste de mots du dictionnaire pour tester les répertoires.
Affiche les résultats en temps réel pendant le scan.
Fournit des options pour ajuster le comportement du scan.

## Auteur
Jordan Munoz - https://github.com/Jordan-Munoz

## Licence
Ce projet est sous licence MIT - voir le fichier LICENSE pour plus de détails.