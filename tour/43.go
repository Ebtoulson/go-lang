// Implement a fibonacci function that returns a
// function (a closure) that returns successive fibonacci numbers.

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    var prev, current = 0, 0
    current = 0
    return func() int{
        if prev == 0 {
            prev = 1
            return current
        }
        prev, current = current, current+prev
        return current
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}