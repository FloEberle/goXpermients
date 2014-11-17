package main

import (
    "fmt"
    "math/big"
    )


// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() *big.Int {  
    i := big.NewInt(0) 
    a := big.NewInt(0)   
    b := big.NewInt(1)   
    return func() *big.Int {  
        i = a  
        a = a.Add(a, b)
        b = i  
        return i  
    }  
} 

func main() {
    f := fibonacci()
    for i := 0; i < 10000; i++ {
        fmt.Println(f())
    }
}