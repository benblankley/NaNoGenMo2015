package main

 import (
     "os"
     "io"
     "fmt"
     "strconv"
     "encoding/json"
     "io/ioutil"
     "log"
)

type AdjectivesType struct {
	Description string `json:"description"`
	Adjs []string `json:"adjs"`
}

type AppliancesType struct {
    Description string `json:"description"`
    Appliances []string `json:"appliances"`
}

type FruitsType struct {
    Description string `json:"description"`
    Fruits []string `json:"fruits"`
}

type ResumeActionWordsType struct {
    Description string `json:"description"`
    Source string `json:"source"`
    ResumeActionWords []string `json:"resume_action_words"`
}

func main() {

    jsonStream, err := ioutil.ReadFile("./adjs.json")
    if err != nil {
        log.Fatal(err)
    }

    var m AdjectivesType
    if err := json.Unmarshal(jsonStream, &m); err != nil {
        log.Fatal(err)
    }

    // Print out entire Adjectives slice
    // fmt.Printf("%s\n", m.Adjs)
    
    // Print length of adjectives slice
    fmt.Printf("Length of Adjectives slice: %d \n", len(m.Adjs))
    
    // Print each item of the Adjectives slice
    // for v := range m.Adjs {
    //	fmt.Println(m.Adjs[v])
    //    }
    

    filename := "output.txt"
    
    file, err := os.Create(filename)
    
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println("Writing to file : " + filename)

    for poemnum := 0; poemnum < 1; poemnum++ {
        fmt.Println("**Poem #" + strconv.Itoa(poemnum))

        io.WriteString(file, "\n**Poem #" + strconv.Itoa(poemnum) + "\n")
        
        io.WriteString(file, "I have eaten\n")
        io.WriteString(file, "the plums\n")
        io.WriteString(file, "that were in\n")
        io.WriteString(file, "the icebox\n\n")
        
        io.WriteString(file, "and which\n")
        io.WriteString(file, "you were probably\n")
        io.WriteString(file, "saving\n")
        io.WriteString(file, "for breakfast\n\n")
        
        io.WriteString(file, "Forgive me\n")
        io.WriteString(file, "they were delicious\n")
        io.WriteString(file, "so sweet\n")
        io.WriteString(file, "and so cold\n")
        
    }
    
    file.Close()
}