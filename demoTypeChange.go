package main 

import (
	"fmt"
)

func runTypeChange() {
   var sum int = 90
   var count int = 5
   var mean float32
   var s string
   
   mean = float32(sum)/float32(count)
   fmt.Printf("Value of mean : %f\n",mean)
   
   s = string(sum)
   fmt.Println(s)
}

func init(){
//	runTypeChange()
}
