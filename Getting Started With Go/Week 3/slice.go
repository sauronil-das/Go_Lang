package main

import(
    "fmt"
    "sort"
    "strconv"
)

func main(){
    var s []int
    var x string
    fmt.Println("Press X to exit, Enter Interger Value: ")
    fmt.Scanln(&x)
    for x!="X" {
        a, err :=strconv.Atoi(x)
        if err == nil {
          s=append(s,a)
          sort.Ints(s)
          fmt.Println(s)
          fmt.Scanln(&x)
        }
    }
}