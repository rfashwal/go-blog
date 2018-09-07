package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rabih/go-blog/accountservice/dbclient"
	"github.com/rabih/go-blog/accountservice/model"
)

var (
	DBClient dbclient.IBoltClient
	client   = &http.Client{}
)

func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
}

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {

	var accountId = mux.Vars(r)["accountId"]

	account, err := DBClient.QueryAccount(accountId)
	account.ServedBy = getIP()
	quote, err := getQuote()
	if err == nil {
		account.Quote = quote
	}

	data, _ := json.Marshal(account)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getQuote() (model.Quote, error) {
	req, _ := http.NewRequest("GET", "http://quotes-service:8080/api/quote?strength=4", nil)
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		quote := model.Quote{}
		bytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bytes, &quote)
		return quote, nil
	} else {
		return model.Quote{}, fmt.Errorf("Some error")
	}
}

// ADD THIS FUNC
func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
	dbUp := DBClient.Check()
	if dbUp {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJsonResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "Database unaccessible"})
		writeJsonResponse(w, http.StatusServiceUnavailable, data)
	}
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
