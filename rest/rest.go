package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/6233/jhcoin/blockchain"
	"github.com/6233/jhcoin/utils"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         URL("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	utils.HandleErr(json.NewEncoder(rw).Encode(data))
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks()))
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.Blockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)

	if err == blockchain.ErrNotFound {
		utils.HandleErr(encoder.Encode(errorResponse{fmt.Sprint(err)}))
	} else {
		utils.HandleErr(encoder.Encode(block))
	}

}

func status(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain()))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	port := fmt.Sprintf(":%d", aPort)

	router := mux.NewRouter()

	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", documentation).Methods("GET")

	router.HandleFunc("/status", status)

	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")

	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}