package main

import (
	"fmt"
    "strconv"
)

func main() {
    for i:= 1; i<=10 ; i++ {
        a := "test"
        b := strconv.Itoa(i)
        fmt.Println(a+b)
    }
  }
  
  //一个可以返回多个值的函数
func numbers()(int,int,string){
    a , b , c := 1 , 2 , "str"
    return a,b,c
  }