# AxoClash Bot Template - Go

Un template de bot pour le jeu AxoClash, implémenté en Go avec le framework Gin.

## Description

Ce projet fournit un bot de base pour AxoClash qui implémente l'API REST requise. Le bot prend des décisions aléatoirement parmi les actions disponibles (déplacements et tirs).

### Fonctionnalités

- API REST conforme à la spécification OpenAPI AxoClash
- Endpoint `/api/v1/bots/info` pour récupérer les informations du bot
- Endpoint `/api/v1/bots/play` pour jouer des actions aléatoirement
- Architecture clean avec séparation models/services/controllers

## Prérequis

- Go 1.27 ou supérieur

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

Le serveur démarre sur le port **9000**.

### Variables d'environnement

| Variable | Description | Valeur par défaut |
|----------|-------------|-------------------|
| `PORT`   | Port d'écoute du serveur | `9000` |

```bash
PORT=8080 go run main.go
```

### Docker

Construire l'image :
```bash
docker build -t axoclash-bot-go .
```

Lancer le conteneur :
```bash
docker run -p 9000:9000 axoclash-bot-go
```

Le conteneur écoute sur le port 9000 par défaut. Le surcharger avec la variable d'environnement `PORT` (et mapper le port en conséquence) :
```bash
docker run -e PORT=8080 -p 8080:8080 axoclash-bot-go
```

L'image finale est construite `FROM scratch` : elle ne contient que le binaire statique (pas d'OS, pas de librairies, pas de certificats — le bot n'appelle aucune API externe).

## API Endpoints

- `GET /api/v1/bots/info` - Retourne les informations du bot (nom et version)
- `POST /api/v1/bots/play` - Reçoit l'état du jeu et retourne une action à jouer

## Structure du projet

```
.
├── controllers/     # Contrôleurs HTTP
├── models/         # DTOs et structures de données  
├── services/       # Logique métier du bot
├── main.go         # Point d'entrée de l'application
├── Dockerfile      # Build de l'image Docker
└── .dockerignore   # Fichiers exclus du contexte de build Docker
```

## Personnalisation

Pour personnaliser le comportement du bot, modifiez la méthode `Play` dans `services/bot_service.go`.
