package main

import (
	"log"

	"github.com/TykTechnologies/tyk-protobuf/bindings/go"
)

// MyPreHook performs a header injection:
func MyPreHook(object *coprocess.Object) (*coprocess.Object, error) {
	object.Request.SetHeaders = map[string]string{
		"Myheader": "Myvalue",
	}
	return object, nil
}

// MyAuthCheck will initialize and return a valid session object if the authentication is ok:
func MyAuthCheck(object *coprocess.Object) (*coprocess.Object, error) {
	authHeader := object.Request.Headers["Authorization"]

	validKey := "d29e8f389a6cf39a72bc7156c5e710885e4b5b89"

	// If the header value doesn't match our "valid key", we don't modify the object:
	if authHeader != validKey {
		log.Println("Bad authentication on MyAuthCheck")
		return object, nil
	}

	log.Println("Successful authentication on MyAuthCheck")

	// If the header value matches, the authentication is correct and we provide a session object:
	object.Session = &coprocess.SessionState{
		Rate: 1000.0,
		Per:  1.0,
	}
	object.Metadata = map[string]string{
		"token": validKey,
	}

	return object, nil

}
