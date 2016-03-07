package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mrjones/oauth"
)

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
	http.Redirect(w, r, redirectURL(r, "/me"), http.StatusTemporaryRedirect)
	return nil
}

func (h *handler) me(w http.ResponseWriter, r *http.Request) error {
	token, err := h.storage.GetAccessToken(r)
	if err != nil {
		return err
	}
	api := anaconda.NewTwitterApi(token.Token, token.Secret)
	user, err := api.GetSelf(url.Values{})
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(user)
}
