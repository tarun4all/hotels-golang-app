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

func GetGeolocation(geolocationService *gl.GeolocationService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestIPAddr, err := getURLParam(r.RequestURI)
		fmt.Println(requestIPAddr, err)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(400) // Return 400 Bad Request.
			return
		}
		info, err := geolocationService.GetGeolocation(requestIPAddr)
		json.NewEncoder(w).Encode(info)
	}
}

func New(geolocationService *gl.GeolocationService) {
	fmt.Println("Setting up router...")
	http.HandleFunc("/hotel/", GetGeolocation(geolocationService))
	log.Fatal(http.ListenAndServe(":3001", nil))
}
