package main

import (
	"encoding/json"
	"fmt"
)

type secretAgent struct {
	First   string `json:"First"`
	Last    string `json:"Last"`
	Age     int    `json:"Age"`
	License bool   `json:"License"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"License":true},{"First":"Alec","Last":"Trevelyan","Age":38,"License":true}]`
	bs := []byte(s)
	fmt.Println(bs)

	var agents []secretAgent

	err := json.Unmarshal(bs, &agents)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(agents)
}
