// Copyright © 2019 Sonoko Mizuki, Ltd. All rights reserved.

package controller

import (
	"encoding/json"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "hello world"})
}

func info(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "alive"})
}

func Run(endpoint string) error {
	http.HandleFunc("/", root)
	http.HandleFunc("/info", info)
	return http.ListenAndServe(endpoint, nil)
}

