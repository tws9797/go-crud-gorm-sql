package main

import (
	"go-crud-gorm-sql/controllers"
	"go-crud-gorm-sql/models"
	"log"
	"net/http"
)

func init() {
	models.Setup()
}

func main() {
	http.HandleFunc("/api/albums", controllers.GetAlbums)
	http.HandleFunc("/api/albums/read", controllers.GetAlbum)
	http.HandleFunc("/api/albums/create", controllers.CreateAlbum)
	http.HandleFunc("/api/albums/update", controllers.UpdateAlbum)
	http.HandleFunc("/api/albums/delete", controllers.DeleteAlbum)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
