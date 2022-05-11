package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(index))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func handleTemp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("%v", cpuTemp)))
		return
	case http.MethodPost:
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		cpuTemp, _ = strconv.ParseFloat(string(body), 64)
		cpuTempGauge.Set(cpuTemp)
		_, _ = w.Write(body)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"msg": "ok"}`))
}

