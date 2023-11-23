package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    moi := Person{"Loïc", 21}
    fmt.Println(SayHello(&moi))
}

func SayHello(p *Person) string {
    return fmt.Sprintf("Hello %s", p.Name)
}
