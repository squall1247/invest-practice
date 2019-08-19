package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		err_msg := "Usage: main arg1 arg2\n arg1 is your API key\n arg2 is stock symbol"
		fmt.Println(err_msg)
		os.Exit(0)
	}

	arg_api_key := strings.Join(os.Args[1:2], "")
	arg_stock := strings.Join(os.Args[2:3], "")
	log.Println(arg_api_key)
	log.Println(arg_stock)

	//Example:
	// https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	// res, err := http.Get("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo")
	req, err := http.NewRequest("GET", "https://www.alphavantage.co/query", nil)

	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("function", "GLOBAL_QUOTE")
	q.Add("symbol", arg_stock)
	q.Add("apikey", arg_api_key)
	req.URL.RawQuery = q.Encode()

	fmt.Println("Send " + req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	resp_json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp_json))
}
