package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

// type Person struct {
// 	Name      string `json:"name"`
// 	Role      string `json:"role"`
// 	Email     string `json:"email"`
// 	Phone     string `json:"phone"`
// 	Snapchat  string `json:"snapchat"`
// 	Instagram string `json:"instagram"`
// }

type facebookAccount struct {
	Email string `json: "EMAIL"`
	Phone string `json: "PHONE"`
	Year  string `json: "DOBY"`
	Month string `json: "DOBM"`
	Day   string `json: "DOBD"`
	Last  string `json: "LN"`
	First string `json: "FN"`
	// State  string `json: ST`
}

func getCSV() {
	resp, err := http.Get("http://nagaganesh.com/actives.csv")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	// var people []Person
	var facebookPayload []facebookAccount
	reader.Read()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if len(line[5]) == 0 {
			continue
		}

		date, err := time.Parse("1/2/2006", line[5])
		name := strings.Split(line[0], " ")
		if err != nil {
			log.Fatal(err)
		}

		// people = append(people, Person{
		// 	Name:      line[0],
		// 	Role:      line[1],
		// 	Email:     line[3],
		// 	Snapchat:  line[6],
		// 	Instagram: line[7],
		// })

		month := strconv.Itoa(int(date.Month()))
		if len(month) < 2 {
			month = "0" + month
		}

		facebookPayload = append(facebookPayload, facebookAccount{
			Email: line[3],
			Phone: line[4],
			Year:  strconv.Itoa(date.Year()),
			Month: month,
			Day:   strconv.Itoa(date.Day()),
			Last:  name[1],
			First: name[0],
		})
	}

	// customAudienceFB, _ := json.Marshal(people)

	// resp, err := http.Post("url", "application/json", bytes.NewBuffer(customAudienceFB))

	facebookJSON, _ := json.Marshal(facebookPayload)
	fmt.Println(string(facebookJSON))

}
