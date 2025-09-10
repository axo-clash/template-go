# AxoClash Bot Template - Go

Un template de bot pour le jeu AxoClash, implémenté en Go avec le framework Gin.

## Description

Ce projet fournit un bot de base pour AxoClash qui implémente l'API REST requise. Le bot prend des décisions aléatoirement parmi les actions disponibles (déplacements et tirs).

### Fonctionnalités

- API REST conforme à la spécification OpenAPI AxoClash
- Endpoint `/info` pour récupérer les informations du bot
- Endpoint `/play` pour jouer des actions aléatoirement
- Architecture clean avec séparation models/services/controllers

## Prérequis

- Go 1.25 ou supérieur

## Installation et lancement

### Lancer en mode développement
```bash
go run main.go
```

### Build et lancer l'exécutable
```bash
go build -o build/axoclash-bot
./build/axoclash-bot
```

Le serveur démarre sur le port **6000**.

## API Endpoints

- `GET /info` - Retourne les informations du bot (nom et version)
- `POST /play` - Reçoit l'état du jeu et retourne une action à jouer

## Structure du projet

```
.
├── controllers/     # Contrôleurs HTTP
├── models/         # DTOs et structures de données  
├── services/       # Logique métier du bot
└── main.go         # Point d'entrée de l'application
```

## Personnalisation

Pour personnaliser le comportement du bot, modifiez la méthode `Play` dans `services/bot_service.go`.