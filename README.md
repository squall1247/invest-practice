 # invest-practice
A application that can get the historical price of specific stock. It get the latest price and volume information of a stock by [Alpha Vantage API](https://www.alphavantage.co/#about).

## Usage
```sh
$ go build main.go
$./main Arg1 Arg2
```
Arg1 is your API key. You can claim your free API key [here](https://www.alphavantage.co/support/#api-key).
Arg2 is the stock symbol you want to query.

## Test environment
Ubuntu 14.04.1

## Todos
Make this application as service that can receive a request from a client.
