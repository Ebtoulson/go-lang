// Implement WordCount. It should return a map of the counts
// of each “word” in the string s. The wc.Test function runs
// a test suite against the provided function and prints
// success or failure.

package main

import (
    "strings"
    "code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    wmap := make(map[string]int, len(words))
    for _, value := range words {
        wmap[value] = wmap[value] + 1 
    }
    return wmap
}

func main() {
    wc.Test(WordCount)
}