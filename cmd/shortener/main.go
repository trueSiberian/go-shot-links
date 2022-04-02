package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var data = make(map[string]string)

func main() {
	// маршрутизация запросов обработчику
	http.HandleFunc("/", LinkHandler)
	// запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		shotUrl := addLink(string(b))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://localhost:8080/" + shotUrl))
		return
	} else if r.Method == http.MethodGet {
		url := strings.Split(r.URL.Path, "/")
		id := url[1]
		if len(url) == 2 {
			val, ok := data[id]
			if ok {
				w.Header().Set("Location", val)
				w.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
		}

	}

	w.WriteHeader(http.StatusBadRequest)
	return

}

func addLink(url string) string {
	shotUrl := generate(5)
	data[shotUrl] = string(url)
	return shotUrl
}

func generate(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
