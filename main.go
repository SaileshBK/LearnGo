package main

import (
    "fmt"
    "LearnGo/Basics"
)

func main() {
    // Functions
    fmt.Println("Adding two numbers")
    response := basic.AddTwoNumbers(1,2)
    fmt.Println("1+2 =", response)
    
    // Loops
    fmt.Println("\nLooping")
    basic.LoopingFunction()     
}
