package environment

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func ExportEnv()  {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("could't load .env file")
		os.Exit(1)
	}
}
