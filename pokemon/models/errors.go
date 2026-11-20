package models

import "errors"

var (
	ErrNotFound = errors.New("404 - ressource introuvable")
	ErrBadRequest = errors.New("400 - problème client")
	ErrTimeout = errors.New("408 - délai d'attente dépassé")
	Errunavailable = errors.New("503 - le serveur est temporarirement incapable de gérer votre requête")
)
