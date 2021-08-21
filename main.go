package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	const (
		// AnswerYes is Yes
		AnswerYes = "yes"
		// AnswerNo is No
		AnswerNo = "no"
		// AnswerWin is I win
		AnswerWin = "I win!!"
		// AnswerLose is I lose
		AnswerLose = "I lose!!"
	)

	// Question struct to map to the Yes and No responses
	type Question struct {
		Yes string `json:"Yes"`
		No  string `json:"No"`
	}

	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into StartingData
	var StartingData map[string]Question
	err = json.Unmarshal(content, &StartingData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	//log.Printf("StartingData: %+v\n", StartingData)

	pq := "Does it have stripes?"
	qNewAnimal := "Is it a tiger?"
	qDistinctive := "Is it a predator?"

	var newQ Question
	newQ.Yes = qDistinctive
	newQ.No = StartingData[pq].No

	StartingData[qDistinctive] = Question{
		Yes: qNewAnimal,
		No:  StartingData[pq].Yes,
	}

	StartingData[pq] = newQ
	StartingData[qNewAnimal] = Question{
		Yes: "AnswerWin",
		No:  "AnswerLose",
	}

	file, _ := json.MarshalIndent(StartingData, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
