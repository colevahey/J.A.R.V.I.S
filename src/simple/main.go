package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Received holds received json data from wolfram
type Received struct {
	QueryResult struct {
		Success bool
		Numpods int
		Pods    []struct {
			Subpods []struct {
				Plaintext string
			}
		}
	}
}

func input() (text string) {
	text = ""
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	return
}

func welcome() {
	fmt.Println("\033[H\033[2JWelcome to J.A.R.V.I.S")
	fmt.Println("What would like to know?")
	answer := input()
	resp, err := http.Get("https://api.wolframalpha.com/v2/query?appid=5KWG7E-HJEU5JGQX6&output=json&includepodid=Result&input=" + strings.Replace(answer, " ", "%20", -1))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(contents))

	data := Received{}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		fmt.Println(err)
	}
	if data.QueryResult.Success == true && data.QueryResult.Numpods > 0 {
		fmt.Println(data.QueryResult.Pods[0].Subpods[0].Plaintext)
	} else {
		fmt.Println("Query is too complicated for JARVIS V1")
	}
}

func main() {
	welcome()
}
