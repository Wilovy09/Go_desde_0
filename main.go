package main

import "fmt"

func main() {

    // como for
    for i := 0; i < 10; i++{        
        if i == 8 {
            fmt.Println("break")
            break
        } else {
            fmt.Println("Continue", i)
            continue
        }
    }

}