package main

import (
	"net/http"
	"fmt"

	"bufio"
    "encoding/csv"
    "encoding/json"
    "io"
    "log"
    "os"
)

type Person struct {
    Name string   `json:"name"`
	Role  string   `json:"role"`
	Email string  `json:"email"`
	Phone string  `json:"phone"`
	Snapchat string  `json:"snapchat"`
	Instagram string  `json:"instagram"`
    Address   *Address `json:"address,omitempty"`
}

func getCSV() {
	resp, err := http.Get("nagaganesh.com/actives.csv")
	if err != nil {
		return
	}
	defer resp.Body.Close()


	csvFile, _ := os.Open(resp)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var people []Person
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        people = append(people, Person{
            Firstname: line[0],
            Lastname:  line[1],
            Address: &Address{
                City:  line[2],
                State: line[3],
            },
        })
    }
    peopleJson, _ := json.Marshal(people)
    fmt.Println(string(peopleJson))


}