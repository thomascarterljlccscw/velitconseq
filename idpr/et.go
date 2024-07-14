package main

import "fmt"

func identity[T any](arg T) T {
    return arg
}

func main() {
    fmt.Println(identity(5))           // Outputs: 5
    fmt.Println(identity("hello"))     // Outputs: hello
    fmt.Println(identity([]int{1, 2, 3})) // Outputs: [1 2 3]
}
