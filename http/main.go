package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:3000/skills.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// bodyString := string(body)
	// fmt.Println(bodyString)

	// One line
	io.Copy(os.Stdout, resp.Body)
}
