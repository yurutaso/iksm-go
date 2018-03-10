package main

import (
	"fmt"
	"github.com/yurutaso/iksm"
	"log"
	"os"
)

const (
	DBNAME string = `iksm.db`
)

func main() {
	client, err := iksm.Client(os.Getenv(`IKSM_SESSION`))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.UpdateResults(DBNAME, false); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finish")
}
