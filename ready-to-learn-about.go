package main

import (
    "fmt"
    "math/rand"
    "math"
)

var palm, python, tiger bool
var rock, scissors int = 42, 55

func addFunc(x, y, z int) int {
    return x + y + z
}

func main() {
    fmt.Println("Hello World!")
    fmt.Println("Random ", rand.Intn(10))
    fmt.Println(math.Pi)
    fmt.Println("Total: ", addFunc(3,4,5))

    var i int
    fmt.Println(i, palm, python, tiger)

    var sun, moon = "eclipse", "waxing"

    fmt.Println("The value of the var rock is:", rock)                    // "The value of the var rock is: 42"
    fmt.Println("The value of the var scissors is:", scissors)            // "The value of the var scissors is: 55"

    fmt.Println("Look at the moon when it is", moon + ".")                // Concatenating strings - "Look at the moon when it is waxing."
    fmt.Println("A total", sun, "of the Sol.")
}

