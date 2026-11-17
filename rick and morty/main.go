package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://rickandmortyapi.com/api/character"

type Character struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Origin   Origin   `json:"origin"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
}

type Origin struct {
	Name string `json:"name"`
}

type Location struct {
	Name string `json:"name"`
}

var (
	ErrRequestCreation = errors.New("impossible de créer la requête")
	ErrRequestSend     = errors.New("impossible d'envoyer la requête")
	ErrNotFound        = errors.New("ressource introuvable")
	ErrInvalidID       = errors.New("identifiant non valide")
	ErrUnexpectedState = errors.New("réponse inattendue de l'API")
	ErrDecode          = errors.New("impossible de décoder la réponse")
)

func getCharacter(client *http.Client, id string) (*Character, error) {
	endpoint := fmt.Sprintf("%s/%v", baseURL, id)

	// Préparation de la requête GET.
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}

	// Envoi de la requête.
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrRequestSend
	}
	defer response.Body.Close()

	// Vérification du code de statut HTTP.
	switch response.StatusCode {
	case http.StatusOK:
		// Succès : on continue vers le décodage.
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusInternalServerError:
		return nil, ErrInvalidID
	default:
		return nil, ErrUnexpectedState
	}

	// Décodage du JSON dans la structure Character.
	var character Character
	if err := json.NewDecoder(response.Body).Decode(&character); err != nil {
		return nil, ErrDecode
	}

	return &character, nil
}

func main() {
	// Client HTTP avec un timeout pour éviter de bloquer indéfiniment.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	id := "1"

	character, err := getCharacter(client, id)
	if err != nil {
		// Gestion différente selon l'erreur retournée.
		switch {
		case errors.Is(err, ErrNotFound):
			log.Fatalf("Personnage %q introuvable, vérifiez l'id.", id)
		case errors.Is(err, ErrInvalidID):
			log.Fatalf("L'id %q n'est pas valide (un nombre est attendu).", id)
		case errors.Is(err, ErrRequestCreation):
			log.Fatalf("Problème interne lors de la création de la requête.")
		case errors.Is(err, ErrRequestSend):
			log.Fatalf("Impossible de joindre l'API, vérifiez votre connexion.")
		case errors.Is(err, ErrDecode):
			log.Fatalf("La réponse de l'API est illisible.")
		case errors.Is(err, ErrUnexpectedState):
			log.Fatalf("L'API a renvoyé une réponse inattendue.")
		default:
			log.Fatalf("Erreur inconnue : %v", err)
		}
	}

	fmt.Println(*character)
}
