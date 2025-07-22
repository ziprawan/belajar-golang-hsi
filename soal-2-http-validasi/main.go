package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func sendResponseWithCode(w http.ResponseWriter, statusCode int, msg []byte) (int, error) {
	w.WriteHeader(statusCode)
	return w.Write(msg)
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
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

		var errors []string
		query := r.URL.Query()
		status := http.StatusOK

		email := strings.TrimSpace(query.Get("email"))
		ageStr := strings.TrimSpace(query.Get("age"))
		if len(email) == 0 {
			errors = append(errors, "Field email tidak boleh kosong")
		}
		if len(ageStr) == 0 {
			errors = append(errors, "Field age tidak boleh kosong")
		}
		if !isEmailValid(email) {
			errors = append(errors, "Field email bukan email yang valid")
		}

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			errors = append(errors, "Field age bukan angka yang valid")
		}
		if age < 18 {
			errors = append(errors, "Field age tidak valid (minimal bernilai 18)")
		}

		logrus.Infof("Didapat email=%s dan age=%d", email, age)
		logrus.Infof("Permintaan terdapat %d kesalahan kueri", len(errors))

		if len(errors) > 0 {
			response["errors"] = errors
			status = http.StatusBadRequest
		} else {
			response["status"] = "ok"
		}

		marshal, err := json.Marshal(response)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")

			sendResponseWithCode(w, http.StatusInternalServerError, []byte("{\"error\":\"Internal Server Error\"}"))
			return
		}

		sendResponseWithCode(w, status, marshal)
	})

	http.ListenAndServe(":8080", nil)
}
