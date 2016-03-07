package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/mrjones/oauth"
)

const (
	aTokenKey = "accesskey"
	rTokenKey = "requestkey"
)

// Storage is an abstract interface for a pluggable storage backend. It needs to
// be able to store and retrieve user's Twitter OAuth request and access tokens.
type Storage interface {
	GetRequestToken(r *http.Request) (*oauth.RequestToken, error)
	SetRequestToken(w http.ResponseWriter, token *oauth.RequestToken) error

	GetAccessToken(r *http.Request) (*oauth.AccessToken, error)
	SetAccessToken(w http.ResponseWriter, token *oauth.AccessToken) error
}

type secureCookieStorage struct {
	cutter *securecookie.SecureCookie
}

// SecureCookieStorage returns a new implementation of the Storage interface
// where the storage vehicle is actually the user's own browser.
func SecureCookieStorage(hashKey, blockKey []byte) Storage {
	return &secureCookieStorage{securecookie.New(hashKey, blockKey)}
}

func (s *secureCookieStorage) GetRequestToken(r *http.Request) (*oauth.RequestToken, error) {
	ret := new(oauth.RequestToken)
	return ret, s.getCookie(r, rTokenKey, ret)
}

func (s *secureCookieStorage) SetRequestToken(w http.ResponseWriter, token *oauth.RequestToken) error {
	return s.setCookie(w, rTokenKey, token)
}

func (s *secureCookieStorage) GetAccessToken(r *http.Request) (*oauth.AccessToken, error) {
	ret := new(oauth.AccessToken)
	return ret, s.getCookie(r, aTokenKey, ret)
}

func (s *secureCookieStorage) SetAccessToken(w http.ResponseWriter, token *oauth.AccessToken) error {
	return s.setCookie(w, aTokenKey, token)
}

func (s *secureCookieStorage) getCookie(r *http.Request, key string, ret interface{}) error {
	cookie, err := r.Cookie(key)
	if err != nil {
		return err
	}
	return s.cutter.Decode(key, cookie.Value, ret)
}

func (s *secureCookieStorage) setCookie(w http.ResponseWriter, key string, value interface{}) error {
	encoded, err := s.cutter.Encode(key, value)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{Name: key, Value: encoded, Path: "/"})
	return nil
}
