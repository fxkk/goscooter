# goscooter

Client-Cli and go package for [howmuchisthe.fish](http://howmuchisthe.fish/) [[github](https://github.com/theimpossibleastronaut/howmuchisthe.fish)]


## Installation

The latest version can be downloaded from [releases](https://github.com/fxkk/goscooter/releases)

Alternatively you can install the binary with go install.

```bash
$ go install github.com/fxkk/goscooter@latest
```

## Usage

### Package

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


### CLI


```bash
$ ./goscooter -h
A CLI based on the https://howmuchisthe.fish API

Usage:
  goscooter [command]

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

<<<<<<< HEAD
Use "goscooter [command] --help" for more information about a command.
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>


## Acknowledgments

* [howmuchisthe.fish](https://howmuchisthe.fish/)
* [huwmuchisthe.fish (github)](https://github.com/theimpossibleastronaut/howmuchisthe.fish)
* [Credits go to theimpossibleastronaut](https://github.com/theimpossibleastronaut)
=======
Use "scooter [command] --help" for more information about a command.
```
>>>>>>> bab6b56 (Update Readme.md)
