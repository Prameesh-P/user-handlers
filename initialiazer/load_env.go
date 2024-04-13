package initialiazer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	if err := godotenv.Load();err != nil{
		log.Fatalf("fialed to load env")
	}
}



