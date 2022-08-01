package main

import (
	"fmt"
	"log"
	"net/http"
	"personas/controller"
	"personas/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	//Persona
	router.HandleFunc("/persona", controller.CreatePersona).Methods(http.MethodPost)
	router.HandleFunc("/persona", controller.GetPersona).Methods(http.MethodGet)
	router.HandleFunc("/persona", controller.UpdatePersona).Methods(http.MethodPut)
	router.HandleFunc("/persona/{identificador}", controller.DeletePersona).Methods(http.MethodDelete)

	router.HandleFunc("/api/persona/html", controller.GetPersonasHTML).Methods(http.MethodGet)

	log.Printf("Servidor escuchando por el puerto %s\n", utils.GetConfig().ApiServerPort)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", utils.GetConfig().ApiServerPort), router))
}

func init() {
	utils.InitConf()
	utils.GetConnection()
}
