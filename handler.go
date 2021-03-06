package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mrjones/oauth"
)

var errNoTwitterAccess = errors.New("missing Twitter Access")

type handler struct {
	consumer *oauth.Consumer
	storage  Storage
}

func redirectURL(r *http.Request, path string) string {
	scheme := "http"
	if r.TLS != nil {
		scheme += "s"
	}
	return fmt.Sprintf("%s://%s%s", scheme, r.Host, path)
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) error {
	redirectTo := redirectURL(r, "/callback")
	token, requestURL, err := h.consumer.GetRequestTokenAndUrl(redirectTo)
	if err != nil {
		return err
	}
	if err := h.storage.SetRequestToken(w, token); err != nil {
		return err
	}
	http.Redirect(w, r, requestURL, http.StatusTemporaryRedirect)
	return nil
}

func (h *handler) callback(w http.ResponseWriter, r *http.Request) error {
	verificationCode := r.URL.Query().Get("oauth_verifier")
	requestToken, err := h.storage.GetRequestToken(r)
	if err != nil {
		return err
	}
	accessToken, err := h.consumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		return err
	}
	if err := h.storage.SetAccessToken(w, accessToken); err != nil {
		return err
	}
	http.Redirect(w, r, redirectURL(r, "/"), http.StatusTemporaryRedirect)
	return nil
}

func (h *handler) followers(w http.ResponseWriter, r *http.Request) error {
	token, err := h.storage.GetAccessToken(r)
	if err != nil {
		return errNoTwitterAccess
	}
	twitterAPI := &api{anaconda.NewTwitterApi(token.Token, token.Secret)}
	response, err := twitterAPI.earliestFollowers(3)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response)
}
