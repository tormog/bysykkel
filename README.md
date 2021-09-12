# BySykkel API Test

This sample project builds a simple program to interact with BySykkel API. 
The program will show a summary of all stations with corresponding bikes and docks available. 
For information see [API documentation](https://oslobysykkel.no/apne-data/sanntid).

## Build and Run

You can run this project locally if you have [GO installed](https://golang.org/doc/install).
If not, you can use the [Docker approach](https://docs.docker.com/get-docker/). 

### Local

To test
`go test .`
 
To build
`go build -o bysykkel .`

To run
`go run .`
OR (after build)
`./bysykkel`

### Docker

To test and build `docker build . -t bysykkel --progress=plain --no-cache`
Using `--no-cache` will always show stdout from commands within Dockerfile.

To run `docker run bysykkel`