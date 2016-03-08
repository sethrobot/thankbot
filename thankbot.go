package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gorilla/mux"
	"github.com/mrjones/oauth"
)

var (
	consumer = oauth.NewConsumer(
		os.Getenv("TWITTER_KEY"),
		os.Getenv("TWITTER_SECRET"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
)

type fallibleHandler func(w http.ResponseWriter, r *http.Request) error

func catchError(fn fallibleHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			fmt.Fprintf(w, "Error for %q: %v", r.URL.String(), err)
		}
	}
}

func main() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_SECRET"))
	storage := SecureCookieStorage(
		[]byte(os.Getenv("COOKIE_HASH_KEY")),
		[]byte(os.Getenv("COOKIE_BLOCK_KEY")),
	)
	h := &handler{consumer: consumer, storage: storage}
	r := mux.NewRouter()
	r.HandleFunc("/login", catchError(h.login)).Methods("GET")
	r.HandleFunc("/callback", catchError(h.callback)).Methods("GET")
	r.HandleFunc("/followers", catchError(h.followers)).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static"))).Methods("GET")
	log.Printf("Listening on %s", os.Getenv("SERVER_BINDING"))
	http.ListenAndServe(os.Getenv("SERVER_BINDING"), r)
}
