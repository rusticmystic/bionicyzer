# bionicyzer

**A** **comman**dline **to**ol **whi**ch **conv**erts **plain**text **in**to **bio**nic **read**ing **rea**dy **mark**down **tex**t. **Th**e **mark**down **gener**ated **b**y **th**e **to**ol **ca**n **ren**der **bio**nic **read**ing **i**n **vsc**ode **prev**iew, **git**hub **REA**DME **et**c.

It can be used along with `gum` to do bionic reading on terminal.



## Build


Clone the repo: git clone https://github.com/rusticmystic/bionicyzer

Install dependencies: `go mod tidy`

Build the tool: `go build -o bionicyzer`

_Then copy the `bionicyzer` tool to `~/.bin` and set the path to that folder. `export PATH=$PATH:~/.bin/`_

## Usage

The `bionicyzer` tool reads from `stdin` and outputs to `stdin`.


To run the tool, use the command: 

```sh

echo "This is an example" | bionicyzer > bionic-out.md

```

We can convert a plain text file into binoinc ready markdown file

```sh
cat file.txt | bionicyzer > output.md

```


## License
This tool is licensed under the MIT License - see the LICENSE file for details.

## Contributing
To contribute to this tool, please fork this repository and create a pull request with your changes.
