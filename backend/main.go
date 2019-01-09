package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Address struct {
	Line1  string `json:"line1"`
	Line2  string `json:"lind2"`
	City   string `json:"city"`
	State  string `json:"state"`
	County string `json:"county"`
	Zip1   string `json:"zip"`
}

type Vet struct {
	LastName      string  `json:"lastName"`
	FirstName     string  `json:"firstName"`
	MiddleName    string  `json:"middleName"`
	FormattedName string  `json:"formattedName"`
	Address       Address `json:"address"`
}

//TODO: ensure vetsByState is a singleton
var vetsByState map[string][]Vet

func getVetsByState(file string) map[string][]Vet {
	if vetsByState == nil {
		vetsByState = make(map[string][]Vet)
	}

	csvFile, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//sanitize
		if strings.ToLower(line[20]) == "active" && line[7] != "" && line[9] != "" && line[10] != "" {
			vetsByState[line[10]] = append(vetsByState[line[10]], Vet{
				LastName:      line[0],
				FirstName:     line[1],
				MiddleName:    line[2],
				FormattedName: line[5],
				Address: Address{
					Line1:  line[7],
					Line2:  line[8],
					City:   line[9],
					State:  line[10],
					County: line[11],
					Zip1:   line[12],
				},
			})
		}
	}

	return vetsByState
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(vetsByState["CO"])
}

// func GetByState(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	_ = getVetsByState("vets2018.csv")
	router := mux.NewRouter()
	router.HandleFunc("/getAll", GetAll).Methods("GET")
	// router.HandleFunc("getByState/{state}", GetByState).methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
