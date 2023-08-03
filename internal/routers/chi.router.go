package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jeanSagaz/go-api/internal/album/domain/entity"
	"github.com/jeanSagaz/go-api/pkg/generics"
)

func ChiHandleRequests() {
	fmt.Println("Rest API v2.0 - chi Routers")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/albums", Router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Router(r chi.Router) {
	r.Get("/", getAlbumsChi)
	r.Get("/{id}", getAlbumByIdChi)
	r.Post("/", postAlbumsChi)
	r.Put("/{id}", putAlbumChi)
	r.Delete("/{id}", deleteAlbumChi)
}

func getAlbumsChi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAlbumsChi")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func getAlbumByIdChi(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.Id == id {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func postAlbumsChi(w http.ResponseWriter, r *http.Request) {
	var newAlbum entity.Album

	// Call ioutil to bind the received JSON to
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(reqBody, &newAlbum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAlbum)
}

func putAlbumChi(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedAlbum entity.Album

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	// al := generics.FirstOrDefault(albums, func(p *models.Album) bool { return p.Id == id })
	// if al == nil {
	// 	result := map[string]any{
	// 		"error":   true,
	// 		"message": "album not found",
	// 	}

	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(result)
	// 	return
	// }

	idx := generics.Find(albums, func(value interface{}) bool {
		return value.(entity.Album).Id == id
	})
	if idx < 0 {
		result := map[string]any{
			"error":   true,
			"message": "album not found",
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
		return
	}

	json.Unmarshal(reqBody, &updatedAlbum)
	albums[idx] = updatedAlbum

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result := map[string]any{
		"error":   false,
		"message": "album updated successfully",
		"data":    updatedAlbum,
	}
	//json.NewEncoder(w).Encode(updatedAlbum)
	json.NewEncoder(w).Encode(result)
}

func deleteAlbumChi(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	al := generics.FirstOrDefault(albums, func(p *entity.Album) bool { return p.Id == id })
	if al == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println(al)

	idx := generics.Find(albums, func(value interface{}) bool {
		return value.(entity.Album).Id == id
	})
	if idx < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//albums = RemoveIndex(albums, idx)
	albums = generics.FindAndDelete(albums, func(p *entity.Album) bool { return p.Id == id })

	fmt.Println(albums)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
