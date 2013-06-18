// Find the cube root of 2, just to make sure the 
// algorithm works. There is a Pow function in 
// the math/cmplx package.

package main

import (
    "math/cmplx"
    "fmt"
)

func Cbrt(x complex128) complex128 {
    z := complex128(1)
    avg := complex128(0)
    for {
        z = z - (cmplx.Pow(z, 3) - x)/(3 * cmplx.Pow(z, 2))
        if cmplx.Abs(z-avg) < 1e-15 {
            break
        }
        avg = z
    }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(Cbrt(8))
}