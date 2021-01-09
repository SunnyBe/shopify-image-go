package router

import (
	"fmt"
	"net/http"
)

func RouterTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an Image upload project.")
}
