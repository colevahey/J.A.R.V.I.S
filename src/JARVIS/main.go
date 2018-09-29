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
		Success   bool
		Numpods   int
		Datatypes string
		Pods      []struct {
			Title      string
			Numsubpods int
			ID         string
			Subpods    []struct {
				Title     string
				Plaintext string
				Img       struct {
					Src string
					Alt string
				}
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
	answer := input()
	resp, err := http.Get("https://api.wolframalpha.com/v2/query?appid=5KWG7E-HJEU5JGQX6&output=json&input=" + strings.Replace(answer, " ", "%20", -1))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(contents))

	data := Received{}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		fmt.Println(err)
	}
	if data.QueryResult.Success {
		if data.QueryResult.Numpods > 0 {
			for _, pod := range data.QueryResult.Pods {
				fmt.Println(pod.Title)
				if strings.Contains(pod.Title, "Result") {
					if pod.Subpods[0].Plaintext != "" {
						fmt.Println("Main result:", pod.Subpods[0].Plaintext)
						fmt.Println(pod.Subpods[0].Img.Src)
					} else {
						fmt.Println("Main result is not plaintext.")
						fmt.Println("Results:", pod.Subpods[0].Img.Src)
					}
				} else {
					if pod.Subpods[0].Plaintext != "" {
						fmt.Println(pod.Subpods[0].Plaintext)
						fmt.Println()
					} else {
						fmt.Println(pod.Subpods[0].Img.Src)
					}
				}
			}
		} else {
			fmt.Println("No relevant results")
		}
	} else {
		fmt.Println("Query was unsuccessful")
	}
}

func main() {
	welcome()
}
