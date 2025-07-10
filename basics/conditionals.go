package basic

import "fmt"

type Options struct {
    Value int
    UseSwitch bool
}

func IfElseSwitch(x Options) {
    value := x.Value
    if x.UseSwitch {
        UseSwitch(value)
        return
    }

    if value%2 == 0 {
        fmt.Printf("%d is an Even.\n", value)
    } else {
        fmt.Printf("%d is an Odd.\n", value)
    }
}

func UseSwitch(x int) {
    switch {
        case x%2 == 0:
            fmt.Printf("Using Switch, %d is an Even.\n", x)
        default:
            fmt.Printf("Using Switch, %d is an Odd.", x)
    }
}