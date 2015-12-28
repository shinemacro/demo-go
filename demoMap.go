package main 

import (
	"fmt"
)

func runMap() {
	/* create a map by make*/
	demoMap := make(map[string]string)
	
	/* insert key-value pairs in the map*/
	demoMap["name"] = "shinemacro"
	demoMap["sex"] = "male"
	demoMap["age"] = "30"
	
	for key, value := range demoMap{
		fmt.Println(key,demoMap[key],value)
	}
	
	/* test if entry is present in the map or not*/
	value, ok := demoMap["age"]
	if(ok){
		fmt.Println("get demoMap key age:" + value)
	}else{
		fmt.Println("demoMap key age isn't exsit.")
	}
	
	/* delete an entry */
	delete(demoMap,"age")
   	fmt.Println("map delete an entry, Entry for age is deleted")  
   	
   	value, ok = demoMap["age"]
	if(ok){
		fmt.Println("get demoMap key age:" + value)
	}else{
		fmt.Println("demoMap key age isn't exsit.")
	}
	
	for key, value := range demoMap{
		fmt.Println(key,demoMap[key],value)
	}
}

func init(){
//	runMap()
}
 