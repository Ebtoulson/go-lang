// Implement Pic. It should return a slice of length dy,
// each element of which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture,
// interpreting the integers as grayscale (well, bluescale) values.

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    picture := make([][]uint8, dy)
    for y := range picture {
        picture[y] = make([]uint8, dx)
        for x := range picture[y] {
            picture[y][x] = uint8(x*y/256)
        } 
    }
    return picture
}

func main() {
    pic.Show(Pic)
}