package auth

import (
	"golang.org/x/oauth2/google"
	"net/http"
	"io/ioutil"
	"log"
	"golang.org/x/net/context"
	"os"
)

var scopes = []string{"https://www.googleapis.com/auth/calendar"}
var key []byte = nil

func readKey() []byte {
	var key, err = ioutil.ReadFile("/tmp/" + os.Getenv("GOOGLE_API_KEY_FILE"))
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
	return googleOauthConfig.Client(context.TODO())
}
