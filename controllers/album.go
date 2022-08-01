package controllers

import (
	"encoding/json"
	"go-crud-gorm-sql/models"
	"log"
	"net/http"
)

var err error

type response struct {
	ID      uint        `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var albums []models.Album
	models.AlbumIndex(&albums)

	if err := json.NewEncoder(w).Encode(albums); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var album models.Album
	var id uint

	if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		log.Fatal(err)
	}

	id = album.ID
	album = models.AlbumByID(id)

	if err := json.NewEncoder(w).Encode(album); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var album models.Album

		if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
			http.Error(w, "Error decoding request body", http.StatusInternalServerError)
			log.Fatal(err)
		}

		models.AlbumCreate(&album)

		res := response{
			ID:      album.ID,
			Message: "User created successfully",
			Data:    album,
		}

		if err = json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var album models.Album

		if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
			http.Error(w, "Error decoding request body", http.StatusInternalServerError)
			log.Fatal(err)
		}

		models.AlbumUpdate(album)

		res := response{
			ID:      album.ID,
			Message: "User updated successfully",
			Data:    album,
		}

		if err = json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var album models.Album

		if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
			http.Error(w, "Error decoding request body", http.StatusInternalServerError)
			log.Fatal(err)
		}

		models.AlbumDelete(album.ID)

		res := response{
			ID:      album.ID,
			Message: "User deleted successfully",
			Data:    album,
		}

		if err = json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
