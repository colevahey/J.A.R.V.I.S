package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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
	fmt.Println("What is it you would like to know?")
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
}

func main() {
	welcome()
}
