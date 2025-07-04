package main

import (
    "fmt"
    "LearnGo/Basics"
)


func main() {
    fmt.Println("Adding two numbers")
    response := functions.AddTwoNumbers(1,2)
    fmt.Println("1+2 =", response)
    
    // Loops
    fmt.Println("\nLooping")
    functions.LoopingFunction()     
}
