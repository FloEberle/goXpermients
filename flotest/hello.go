package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"os"
	"math"
)

func main() {
	resp, err := http.Get("http://www.stv-wettingen.ch")
	if err != nil {
		fmt.Printf("%s \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Hello, world.  %s\n", body)
	fmt.Printf("Time: %s \n", time.Now())
	fmt.Printf("%s \n", math.Pi)
	fmt.Println(add(10, 32))
}

func add(x int, y int) int {
	return x + y
}
