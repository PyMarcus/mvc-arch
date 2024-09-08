package main 

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println("Error ", err)
	}
	fmt.Println(os.Getenv("TEST"))
}
