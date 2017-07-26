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

func readKey(path string) []byte {
	var k, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return k
}

func RequireGoogleClient() *http.Client {
	if key == nil {
		key = readKey("/tmp/" + os.Getenv("GOOGLE_API_KEY_FILE"))
	}
	var googleOauthConfig, err = google.JWTConfigFromJSON(key, scopes...)
	if err != nil {
		log.Fatal(err)
	}
	return googleOauthConfig.Client(context.TODO())
}
