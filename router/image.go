package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Image struct {
	Id      string `json:id`
	Title   string `json:"title`
	Url     string `json:"url`
	Content string `json:"content`
}

var Images []Image = SetTestImages()

func SetTestImages() []Image {
	return []Image{
		{
			Id:      "1",
			Title:   "This is a test Image",
			Url:     "image url goes here",
			Content: "image content goes here",
		},

		{
			Id:      "2",
			Title:   "This is a second test Image",
			Url:     "image url goes here",
			Content: "image content goes here",
		},
	}
}

func SelectImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Select all images.")
	json.NewEncoder(w).Encode(Images)
}

func SelectImageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintln(w, "Select image with id: "+key)
	for _, image := range Images {
		if image.Id == key {
			json.NewEncoder(w).Encode(image)
		}
	}
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new image.")

	// Read request body as string
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))

	// Unmarshals the JSON in the request body into a new image
	var image Image
	json.Unmarshal(reqBody, &image)

	Images = append(Images, image)
	json.NewEncoder(w).Encode(image)
}

func UpdateImageById(w http.ResponseWriter, r *http.Request) {
	// Get reqBody json as a string
	reqBody, _ := ioutil.ReadAll(r.Body)
	key := mux.Vars(r)["id"]

	fmt.Fprintln(w, "Updating item with id "+key)

	// Unmarshal the new image content
	var newImage Image
	json.Unmarshal(reqBody, &newImage)

	// If key is found in image list update update that image instance
	for index, image := range Images {
		if image.Id == key {
			*&image = newImage
			*&Images[index] = image
			json.NewEncoder(w).Encode(image)
		}
	}
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]

	// If key is found in image list update update that image instance
	for index, image := range Images {
		if image.Id == key {
			// Removes the image from the list of images
			Images = append(Images[:index], Images[index+1:]...)
			json.NewEncoder(w).Encode(image)
		}
	}
}

// curl -d '{"id":"3", "title":"Added Title", "url":"Added Url", "content":"Added Content",}' -H "Content-Type: application/json" -X POST http://localhost:8001/image/create
// curl -d '{"id":"3", "title":"Updated Title", "url":"Updated Url", "content":"Updated Content"}' -H "Content-Type: application/json" -X POST http://localhost:8001/image/update/1
// curl -d -H "Content-Type: application/json" -X GET http://localhost:8001/image/delete/1
