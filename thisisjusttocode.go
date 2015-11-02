package main

 import (
     "os"
     "io"
     "fmt"
     "strconv"
 )

func main() {
    
    filename := "output.txt"
    
    file, err := os.Create(filename)
    
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println("Writing to file : " + filename)

    for poemnum := 0; poemnum < 18; poemnum++ {
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