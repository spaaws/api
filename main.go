package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Descritpion string    `json:"description"`
}

type allEvents []Event

var events = allEvents{}

// @title           Swagger API-Event
// @version         1.0
// @description     Documentação da API de eventos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Event Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://new-api-events.herokuapp.com/
func main() {
	log.Println("Starting API")
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	router.HandleFunc("/event", CreateNewEvent).Methods("POST")

	http.ListenAndServe(":"+port, router)
	// http.ListenAndServe(":8080", router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação Funcionando...\n")
}

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  Event
// @Router       /events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// CreateNewEvent godoc
// @Summary      New event
// @Description  Creater new events
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id    title    description    "Event"
// @Succees      201  {object}  Event
// @Router       /event [post]
func CreateNewEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint create new event")
	var newEvent Event

	newEvent.Id, _ = uuid.NewV4()

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEvent)

	events = append(events, newEvent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}
