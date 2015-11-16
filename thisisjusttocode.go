package main

import (
	"encoding/json"
	"fmt"
	"github.com/signintech/gopdf"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

type VerbsType struct {
	Description string   `json:"description"`
	Source      string   `json:"source"`
	Verbs       []string `json:"resume_action_words"`
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
	// load fruits json
	jsonStream, err3 := ioutil.ReadFile("./fruits.json")
	if err3 != nil {
		log.Fatal(err3)
	}
	// import into p slice
	var p FruitsType
	if err := json.Unmarshal(jsonStream, &p); err != nil {
		log.Fatal(err)
	}
	// load verbs json
	jsonStream, err4 := ioutil.ReadFile("./resume_action_words.json")
	if err4 != nil {
		log.Fatal(err4)
	}
	// import into q slice
	var q VerbsType
	if err := json.Unmarshal(jsonStream, &q); err != nil {
		log.Fatal(err)
	}

	// Print length of adjectives slice
	fmt.Printf("Length of Adjectives slice: %d \n", len(m.Adjs))
	// Print length of appliances slice
	fmt.Printf("Length of Appliances slice: %d \n", len(n.Appliances))
	// Print length of fruits slice
	fmt.Printf("Length of Fruits slice: %d \n", len(p.Fruits))
	// Print length of verbs slice
	fmt.Printf("Length of Verbs slice: %d \n", len(q.Verbs))

	fmt.Println("Magic adjective 8-Ball says:", m.Adjs[rand.Intn(len(m.Adjs))])
	fmt.Println("Magic appliances 8-Ball says:", strings.ToLower(n.Appliances[rand.Intn(len(n.Appliances))]))
	fmt.Println("Magic fruits 8-Ball says:", p.Fruits[rand.Intn(len(p.Fruits))])
	fmt.Println("Magic verbs 8-Ball says:", q.Verbs[rand.Intn(len(q.Verbs))])

	// Print out entire Adjectives slice
	// fmt.Printf("%s\n", m.Adjs)
	// fmt.Printf("%s\n", n.Appliances)

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

	// add for pdf support
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	err = pdf.AddTTFFont("DejaVuSerif", "./src/github.com/signintech/gopdf/gopdftest/ttf/DejaVuSerif.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	for poemnum := 0; poemnum < 1800; poemnum++ {

		// add for pdf support
		pdf.AddPage()
		pdf.SetLineWidth(2)
		pdf.Line(10, 320, 585, 320)
		err = pdf.SetFont("DejaVuSerif", "", 14)
		if err != nil {
			log.Print(err.Error())
			return
		}

		// output status to screen
		fmt.Println("Poem #" + strconv.Itoa(poemnum+1))

		io.WriteString(file, "\nPoem #"+strconv.Itoa(poemnum+1)+"\n")
		pdf.Cell(nil, "Poem #"+strconv.Itoa(poemnum+1))
		pdf.Br(20)

		s3 := fmt.Sprint("I have ", q.Verbs[rand.Intn(len(q.Verbs))])
		io.WriteString(file, fmt.Sprint(s3, " \n"))
		pdf.Cell(nil, s3)
		pdf.Br(20)

		s0 := fmt.Sprint("the ", p.Fruits[rand.Intn(len(p.Fruits))])
		io.WriteString(file, fmt.Sprint(s0, " \n"))
		pdf.Cell(nil, s0)
		pdf.Br(20)

		io.WriteString(file, "that was in\n")
		pdf.Cell(nil, "that was in")
		pdf.Br(20)

		s1 := fmt.Sprint("the ", strings.ToLower(n.Appliances[rand.Intn(len(n.Appliances))]))
		io.WriteString(file, fmt.Sprint(s1, "\n\n"))
		pdf.Cell(nil, s1)
		pdf.Br(20)
		pdf.Br(20)

		io.WriteString(file, "and which\n")
		pdf.Cell(nil, "and which")
		pdf.Br(20)

		io.WriteString(file, "you were probably\n")
		pdf.Cell(nil, "you were probably")
		pdf.Br(20)

		io.WriteString(file, "saving\n")
		pdf.Cell(nil, "saving")
		pdf.Br(20)

		io.WriteString(file, "for breakfast\n\n")
		pdf.Cell(nil, "for breakfast")
		pdf.Br(20)
		pdf.Br(20)

		io.WriteString(file, "Forgive me\n")
		pdf.Cell(nil, "forgive me")
		pdf.Br(20)

		s2 := fmt.Sprint("it was ", m.Adjs[rand.Intn(len(m.Adjs))])
		io.WriteString(file, fmt.Sprint(s2, " \n"))
		pdf.Cell(nil, s2)
		pdf.Br(20)

		s4 := fmt.Sprint("so ", m.Adjs[rand.Intn(len(m.Adjs))])
		io.WriteString(file, fmt.Sprint(s4, " \n"))
		pdf.Cell(nil, s4)
		pdf.Br(20)

		io.WriteString(file, "and so good\n")
		pdf.Cell(nil, "and so good")
		pdf.Br(20)

	}
	pdf.WritePdf("output.pdf")
	file.Close()
}
