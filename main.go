package main

import (
	"fmt"
	"log"
	"net/http"

	entity "github.com/beilmo/spectre-go-rest-api/entity"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	speaker := &entity.Speaker{
		Id:        101,
		FirstName: "Dorin",
		LastName:  "Danciu",
	}

	data, err := proto.Marshal(speaker)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newSpeaker := &entity.Speaker{}
	err = proto.Unmarshal(data, newSpeaker)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if speaker.GetFirstName() != newSpeaker.GetFirstName() {
		log.Fatalf("data mismatch %q != %q", speaker.GetFirstName(), newSpeaker.GetFirstName())
	}

	fmt.Println("Hello World!")
	i := 101
	fmt.Println(i)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
