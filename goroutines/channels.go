package main

import "fmt"

func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    a := []int{7, 2, 8, -9, 4, 12}

    c := make(chan int)
    d := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], d)
    x, y := <-c, <-d // receive from c

    fmt.Println(x, y, x+y)
}