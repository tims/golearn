package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"log"
)

type Salt struct {
	Text string `json:"text"`
}

func index(w http.ResponseWriter, r *http.Request) {
	salt := Salt{"Around and around the rugged rucks the ragged rascal ran"}
	fmt.Printf("%+v\n", *r)

	// What are you requesting? wat?
	b, err := json.Marshal(salt)
	if err != nil {
		log.Fatal(err);
		io.WriteString(w, "arg")
	} else {
		io.WriteString(w, string(b))
	}
}

func print(c chan int) {
	fmt.Println(<- c);
}

func main() {
	fmt.Println("Hi, I am made of cheese")
	//http.HandleFunc("/", index)
	//http.ListenAndServe(":8000", nil)

	var c = make(chan string);
	for i := 1; i <= 5; i++ {
		go func(i int, co chan<- string) {
			for j := 1; j <= 5; j++ {
				co <- fmt.Sprintf("hi from %d.%d", i, j)
			}
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		fmt.Println(<-c)
	}
}
