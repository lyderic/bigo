package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const file = "/dev/shm/bw.json"

func main() {
	var err error
	var db Database
	var content []byte
	if content, err = os.ReadFile(file); err != nil {
		panic(err)
	}
	fmt.Println("loaded", len(content), "bytes")
	if err = json.Unmarshal(content, &db); err != nil {
		panic(err)
	}
	fmt.Println("decoded", len(db.Items), "items")
	for idx, item := range db.Items {
		fmt.Printf("[%03d] %s\nUsername: %s\nPassword: %s\nURLs: %v\n\n", idx+1, item.Name, item.Login.Username, item.Login.Password, item.Login.Uris)
	}
}
