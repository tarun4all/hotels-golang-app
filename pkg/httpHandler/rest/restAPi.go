package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	gl "github.com/tarun4all/hotels-golang-app/pkg/geolocation"
)

func getQueryParams(url string) map[string]string {
	params := strings.Split(url, "&")
	var paramsMap = make(map[string]string)

	for _, param := range params {
		temp := strings.Split(param, "=")
		paramsMap[temp[0]] = temp[1]
	}

	return paramsMap
}

func getURLParam(url string) (string, error) {
	var re = regexp.MustCompile(`(\?|\&)([^=]+)\=([^&]+)`)

	// clean URL if query params exist
	url = re.ReplaceAllString(url, "")
	urlParams := strings.Split(url, "/")

	fmt.Println(urlParams, len(urlParams))
	if len(urlParams) != 3 || urlParams[2] == "" {
		return "", errors.New("Invalid URL")
	}
	return urlParams[2], nil
}

type HttpHandler struct {
	service *gl.GeolocationService
}

func (handler *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.get(w, r)
}

func (handler *HttpHandler) get(w http.ResponseWriter, r *http.Request) {
	requestIPAddr, err := getURLParam(r.RequestURI)
	fmt.Println(requestIPAddr, err)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}
	info, err := handler.service.GetGeolocation(requestIPAddr)
	json.NewEncoder(w).Encode(info)
}

func NewHandler(service *gl.GeolocationService) http.Handler {
	var handler = HttpHandler{
		service: service,
	}
	return &handler
}
