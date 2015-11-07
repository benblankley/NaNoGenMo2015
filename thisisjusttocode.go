package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type AdjectivesType struct {
	Description string   `json:"description"`
	Adjs        []string `json:"adjs"`
}

type AppliancesType struct {
	Description string   `json:"description"`
	Appliances  []string `json:"appliances"`
}

type FruitsType struct {
	Description string   `json:"description"`
	Fruits      []string `json:"fruits"`
}

type ResumeActionWordsType struct {
	Description       string   `json:"description"`
	Source            string   `json:"source"`
	ResumeActionWords []string `json:"resume_action_words"`
}

func main() {

	// seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// load adjectives json
	jsonStream, err := ioutil.ReadFile("./adjs.json")
	if err != nil {
		log.Fatal(err)
	}
	// import into m slice
	var m AdjectivesType
	if err := json.Unmarshal(jsonStream, &m); err != nil {
		log.Fatal(err)
	}

	// load appliances json
	jsonStream, err2 := ioutil.ReadFile("./appliances.json")
	if err2 != nil {
		log.Fatal(err2)
	}
	// import into n slice
	var n AppliancesType
	if err := json.Unmarshal(jsonStream, &n); err != nil {
		log.Fatal(err)
	}

	// Print out entire Adjectives slice
	// fmt.Printf("%s\n", m.Adjs)
	// fmt.Printf("%s\n", n.Appliances)

	// Print length of adjectives slice
	fmt.Printf("Length of Adjectives slice: %d \n", len(m.Adjs))
	fmt.Println("Magic adjective 8-Ball says:", m.Adjs[rand.Intn(len(m.Adjs))])
	fmt.Println("Magic appliances 8-Ball says:", n.Appliances[rand.Intn(len(n.Appliances))])

	// Print each item of the Adjectives slice
	// for v := range m.Adjs {
	//	fmt.Println(m.Adjs[v])
	// }

	filename := "output.txt"

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Writing to file : " + filename)

	for poemnum := 0; poemnum < 20; poemnum++ {
		fmt.Println("**Poem #" + strconv.Itoa(poemnum+1))

		io.WriteString(file, "\n**Poem #"+strconv.Itoa(poemnum+1)+"\n")

		io.WriteString(file, "I have eaten\n")
		io.WriteString(file, "the plums\n")
		io.WriteString(file, "that were in\n")
		s1 := fmt.Sprint("the ", n.Appliances[rand.Intn(len(n.Appliances))], " \n\n")
		io.WriteString(file, s1)

		io.WriteString(file, "and which\n")
		io.WriteString(file, "you were probably\n")
		io.WriteString(file, "saving\n")
		io.WriteString(file, "for breakfast\n\n")

		io.WriteString(file, "Forgive me\n")
		s2 := fmt.Sprint("they were ", m.Adjs[rand.Intn(len(m.Adjs))], " \n")
		io.WriteString(file, s2)
		io.WriteString(file, "so sweet\n")
		io.WriteString(file, "and so good\n")

	}

	file.Close()
}
