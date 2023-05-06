package main

import "github.com/harshithgowdakt/gopkgexample"
import "fmt"

func main(){

	fmt.Println("This repo is created to understand the go modules");
	
	sum :=gopkgexample.Add(1,2,3);
	fmt.Printf("sum=%d \n", sum);

	product :=gopkgexample.Multiply(2,3,4);
	fmt.Printf("product=%d \n",product);
}