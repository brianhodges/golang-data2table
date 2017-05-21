package main
import (
    "fmt"
    "strings"
    "unicode/utf8"
)

type Country struct {
    Name string
    Population string
}

func main() {
    countries := []Country{
        {
            Name: "USA",
            Population: "350 million",
        },
        {
            Name: "China",
            Population: "1.4 billion",
        },
        {
            Name: "Russia",
            Population: "145 million",
        },
    }
    
    fmt.Println()
    fmt.Println("SANITIZED DUMP:")
	for _, c := range countries {
		fmt.Println("{ :Name => " + c.Name + ", :Population => " + c.Population + " } ")
	}
	
    fmt.Println()
    fmt.Println("TABLE:")
    fmt.Println("Countries | Population")
    fmt.Println(strings.Repeat("-", utf8.RuneCountInString("Countries | Population")))
	for _, c := range countries {
        sp := utf8.RuneCountInString("Countries ") - len(c.Name) 
		fmt.Println(c.Name + strings.Repeat(" ", sp) + "|" + c.Population)
	}
    fmt.Println()
}