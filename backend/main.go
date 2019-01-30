package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Define the Geocode API URL as a constant
const (
	geocodeApiUrl = "https://maps.googleapis.com/maps/api/geocode/json?"
	apiKey        = "AIzaSyCa4g4lRqdn7GDsSDC9lT1VXHpZjdrK470"
)

// All results from the JSON
type Results struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

// Result store each result from the JSON
type Result struct {
	AddressComponents []GAddress `json:"address_components"`
	FormattedAddress  string     `json:"formatted_address"`
	Geometry          Geometry   `json:"geometry"`
	PlaceId           string     `json:"place_id"`
	Types             []string   `json:"types"`
}

// Address store each address is identified by the 'types'
type GAddress struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// Geometry store each value in the geometry
type Geometry struct {
	Bounds       Bounds `json:"bounds"`
	Location     LatLng `json:"location"`
	LocationType string `json:"location_type"`
	Viewport     Bounds `json:"viewport"`
}

// Bounds Northeast and Southwest
type Bounds struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

// LatLng store the latitude and longitude
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Address struct {
	Line1  string `json:"line1"`
	Line2  string `json:"line2"`
	City   string `json:"city"`
	State  string `json:"state"`
	County string `json:"county"`
	Zip1   string `json:"zip"`
}

type Vet struct {
	LastName      string   `json:"lastName"`
	FirstName     string   `json:"firstName"`
	MiddleName    string   `json:"middleName"`
	FormattedName string   `json:"formattedName"`
	Location      Location `json:"location"`
	Address       Address  `json:"address"`
	Distance      float64  `json:"distance"`
}

type ByLocality []Vet

//TODO: ensure vetsByState is a singleton
var vetsByState map[string][]Vet

func (a ByLocality) Len() int           { return len(a) }
func (a ByLocality) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
func (a ByLocality) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
				LastName:      strings.Trim(line[0], " "),
				FirstName:     strings.Trim(line[1], " "),
				MiddleName:    strings.Trim(line[2], " "),
				FormattedName: strings.Trim(line[5], " "),
				Address: Address{
					Line1:  strings.Trim(strings.Replace(line[7], "#", "", -1), " "),
					Line2:  strings.Trim(strings.Replace(line[8], "#", "", -1), " "),
					City:   strings.Trim(line[9], " "),
					State:  strings.Trim(line[10], " "),
					County: strings.Trim(line[11], " "),
					Zip1:   strings.Trim(line[12], " "),
				},
			})
		}
	}

	return vetsByState
}

func (address *Address) GetFormattedAddress() string {
	// Creats a slice with all content from the Address struct
	var content []string

	//remove whitespace and hash signs which may break google geocode api requests
	content = append(content, address.Line1)
	content = append(content, address.Line2)
	content = append(content, address.Zip1)
	content = append(content, address.City)
	content = append(content, address.State)

	var formattedAddress string

	// For each value in the content slice check if it is valid
	// and add to the formattedAddress string
	for _, value := range content {
		if value != "" {
			if formattedAddress != "" {
				formattedAddress += ", "
			}
			formattedAddress += value
		}
	}

	return formattedAddress
}

func httpRequest(url string) (Results, error) {

	var results Results

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}

	// For control over HTTP client headers, redirect policy, and other settings, create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}

	// Callers should close resp.Body when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func appendLocations(vetsByState map[string][]Vet) map[string][]Vet {
	//iterate thru all vets appending lat long

	vetEntry := 0
	for i, vet := range vetsByState["CO"] {
		//Only serve 25 for now so we dont have to parse whole data set..
		//This will be removed when DB logic is added
		if vetEntry == 25 {
			break
		}
		vetEntry++
		log.Printf("on vet entry: %v/%v", vetEntry, len(vetsByState["CO"]))
		//perform google maps api call with address
		fAddr := vet.Address.GetFormattedAddress()
		fAddr = strings.Replace(fAddr, " ", "+", -1)

		url := geocodeApiUrl + "address=" + fAddr

		if apiKey != "" {
			url += "&key=" + apiKey
		}

		//TODO: dispatch these http requests out to a consumer pool of goroutines
		// Send the HTTP request and get the results
		results, err := httpRequest(url)
		if err != nil {
			log.Println(url)
		}

		if len(results.Results) > 0 {
			vetsByState["CO"][i].Location.Latitude = results.Results[0].Geometry.Location.Lat
			vetsByState["CO"][i].Location.Longitude = results.Results[0].Geometry.Location.Lng
		} else {
			log.Printf("Did not get lat lng for: %+v\n", vetsByState["CO"][i])
			log.Println(url)
		}
	}
	vetsByState["CO"] = vetsByState["CO"][:vetEntry]

	return vetsByState
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("PARAMS: %+v\n", r.URL.Query())
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for i, vet := range vetsByState["CO"] {
		lat1, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
		if err != nil {
			log.Println(err)
		}
		lon1, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
		if err != nil {
			log.Println(err)
		}
		vetsByState["CO"][i].Distance = math.Round(100*distance(lat1, lon1, vet.Location.Latitude, vet.Location.Longitude, "M")) / 100
	}

	sort.Sort(ByLocality(vetsByState["CO"]))
	json.NewEncoder(w).Encode(vetsByState["CO"])
}

func main() {
	_ = getVetsByState("vets2018.csv")
	vetsByState = appendLocations(vetsByState)
	router := mux.NewRouter()
	router.HandleFunc("/getAll", GetAll).Methods("GET")
	// router.HandleFunc("getByState/{state}", GetByState).methods("GET")
	log.Println("listen & serve port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
