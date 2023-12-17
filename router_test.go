package gyocharo

/*
import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestGetRoute(t *testing.T) {
	router := NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ur mom"))
	})

	var server http.Server = http.Server{":3000", router}
	log.Fatal(http.ListenAndServe(":3000", router))

	resp, err := http.Get("0.0.0.0:3000/")
	if err != nil {
		t.Errorf("Request not found")
	}
	fmt.Printf("%v\n", resp)

}
*/
