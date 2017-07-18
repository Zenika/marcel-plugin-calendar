package auth

import (
	"golang.org/x/oauth2/google"
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"golang.org/x/net/context"
)

var Google_API_client *http.Client = nil
var defaultNbEvents string = "10"
var scopes = []string{"https://www.googleapis.com/auth/calendar"}
var key []byte = nil

func readKey() []byte {
	var key, err = ioutil.ReadFile(os.Getenv("GOOGLE_API_KEY_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func RequireGoogleClient() *http.Client {
	if key == nil {
		key = readKey()
	}
	var googleOauthConfig, err = google.JWTConfigFromJSON(key, scopes...)
	if err != nil {
		log.Fatal(err)
	}
	return googleOauthConfig.Client(context.Background())
}
