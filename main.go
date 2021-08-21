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
		Yes string // what question is next, if answer is yes
		No  string // what question is next, if answer is no
	}

	//var StartingData map[string]Question

	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var StartingData map[string]interface{}
	err = json.Unmarshal(content, &StartingData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	//log.Printf("origin: %s\n", payload["Does it have 4 legs?"])
	log.Printf("StartingData: %+v\n", StartingData)

	pq := "Does it have stripes?"
	//qNewAnimal := "Is it a tiger?"
	//qDistinctive := "Is it a predator?"
	qPrevious := StartingData[pq]
	log.Printf("%+v", qPrevious)
	log.Print(len(StartingData))

	// StartingData[qDistinctive] = Question{
	// 	Yes: qNewAnimal,
	// 	No:  g.StartingData[pq].Yes,
	// }

	// StartingData[qPrevious] = Question{
	// 	Yes: qDistinctive,
	// }

	// StartingData[pq] = qPrevious

	//log.Printf("StartingData: %+v\n", StartingData)
	file, _ := json.MarshalIndent(StartingData, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
