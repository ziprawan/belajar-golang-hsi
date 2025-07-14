package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func sendResponseWithCode(w http.ResponseWriter, statusCode int, msg []byte) (int, error) {
	w.WriteHeader(statusCode)
	return w.Write(msg)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r == nil {
			logrus.Error("Request is nil")

			sendResponseWithCode(w, http.StatusInternalServerError, []byte("Internal Server Error"))
			return
		}

		logrus.Infof("Terakses %s dengan metode %s", r.URL.Path, r.Method)
		sendResponseWithCode(w, http.StatusOK, []byte("Server sedang berjalan"))
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		if r == nil {
			logrus.Error("Request is nil")

			sendResponseWithCode(w, http.StatusInternalServerError, []byte("Internal Server Error"))
			return
		}

		logrus.Infof("Terakses /validate dengan metode %s", r.Method)

		w.Header().Add("Content-Type", "application/json")

		response := make(map[string]any)
		response["ok"] = true

		marshal, err := json.Marshal(response)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")

			sendResponseWithCode(w, http.StatusInternalServerError, []byte("Internal Server Error"))
			return
		}

		sendResponseWithCode(w, 200, marshal)
	})

	http.ListenAndServe(":8080", nil)
}
