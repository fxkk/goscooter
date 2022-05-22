# goscooter

Client-Cli and go package for [howmuchisthe.fish](http://howmuchisthe.fish/)

## Package

```go
package main

import "github.com/fxkk/goscooter/pgk/scooter"

func main() {

// get Random Quote
quote, _ := scooter.Random()
fmt.println(quote)

// get full response struct
response, _ := scooter.RandomFull()
fmt.Printf("%+v\n", response)
}
```


## CLI


```bash
$ ./scooter -h
A CLI based on the https://howmuchisthe.fish API

Usage:
  scooter [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  daily       Daily Scooter Quote
  help        Help about any command
  quote       Get Quote by id
  random      Random Scooter Quote

Flags:
  -f, --full      full response
  -h, --help      help for scooter
  -v, --verbose   verbose output

Use "scooter [command] --help" for more information about a command.
```
