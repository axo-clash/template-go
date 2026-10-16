# --- Etape 1 : Build de l'application (binaire statique) ---
FROM golang:1.27-alpine AS build
WORKDIR /app

# Copie les fichiers de dependances puis telecharge les modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# CGO desactive + strip des symboles : binaire statique minimal
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /axoclash-bot .

# --- Etape 2 : Image finale minimale (FROM scratch) ---
# Le binaire est statiquement lie : aucune librairie ni certificat
# n'est necessaire au runtime.
FROM scratch

# Port d'ecoute par defaut (surchargeable via la variable d'environnement PORT)
ENV PORT=9000
EXPOSE 9000

COPY --from=build /axoclash-bot /axoclash-bot

ENTRYPOINT ["/axoclash-bot"]
