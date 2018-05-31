package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	. "github.com/abhi06276/contents-restapi/config"
	. "github.com/abhi06276/contents-restapi/dao"
	. "github.com/abhi06276/contents-restapi/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = ContentsDAO{}

func GetAllContents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yetaaa !")
}

func FindContentEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contents, err := dao.FindByAppId(params["app_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid App ID")
		return
	}
	respondWithJson(w, http.StatusOK, contents)
}

func CreateContentEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var content ContentModel

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	content.ID = bson.NewObjectId()
	if err := dao.Insert(content); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, content)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
		fmt.Println("====== Listening to port", port)
		r := mux.NewRouter()
		r.HandleFunc("/content", GetAllContents).Methods("GET")
		r.HandleFunc("/content", CreateContentEndPoint).Methods("POST")
		r.HandleFunc("/content", UpdateMovieEndPoint).Methods("PUT")
		r.HandleFunc("/content", DeleteMovieEndPoint).Methods("DELETE")
		r.HandleFunc("/content/{app_id}", FindContentEndpoint).Methods("GET")
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatal(err)
		}
	}
}
