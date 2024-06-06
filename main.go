package main

import "C"
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/noscrape/noscrape/noscrape"
	"os"
	"seehuhn.de/go/sfnt"
)

type Input struct {
	Font  string           `json:"font"`
	Trans map[string]int32 `json:"translation"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./noscrape '<JSON string>'")
		os.Exit(1)
	}
	var input Input

	// Parse the JSON string into the Input struct
	err := json.Unmarshal([]byte(os.Args[1]), &input)
	if err != nil {
		_ = fmt.Errorf("Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	font, err := sfnt.ReadFile(input.Font)
	if err != nil {
		_ = fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}

	buf, err := noscrape.Render(*font, input.Trans)
	if err != nil {
		_ = fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}

	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Println(b64)

	os.Exit(0)
}
