package main

import (
	"github.com/aktur/imdb-trakt-sync-1/pkg/providers"
	"github.com/aktur/imdb-trakt-sync-1/pkg/providers/imdb"
	"github.com/aktur/imdb-trakt-sync-1/pkg/providers/trakt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	validateEnvVars()
	providers.Sync()
}

func validateEnvVars() {
	_, ok := os.LookupEnv(imdb.ListIdsKey)
	if !ok {
		log.Fatalf("environment variable %s is required", imdb.ListIdsKey)
	}
	_, ok = os.LookupEnv(imdb.CookieAtMainKey)
	if !ok {
		log.Fatalf("environment variable %s is required", imdb.CookieAtMainKey)
	}
	_, ok = os.LookupEnv(imdb.CookieUbidMainKey)
	if !ok {
		log.Fatalf("environment variable %s is required", imdb.CookieUbidMainKey)
	}
	_, ok = os.LookupEnv(trakt.ClientIdKey)
	if !ok {
		log.Fatalf("environment variable %s is required", trakt.ClientIdKey)
	}
	_, ok = os.LookupEnv(trakt.ClientSecretKey)
	if !ok {
		log.Fatalf("environment variable %s is required", trakt.ClientIdKey)
	}
	_, ok = os.LookupEnv(trakt.UsernameKey)
	if !ok {
		log.Fatalf("environment variable %s is required", trakt.UsernameKey)
	}
	_, ok = os.LookupEnv(trakt.PasswordKey)
	if !ok {
		log.Fatalf("environment variable %s is required", trakt.PasswordKey)
	}
}
