# Go4Life

Game of life in golang

## Dev Env

I created a custom Docker container. 

To build it run:

```sh
docker build -t go-4-life .
```

To run it use:

```sh
docker run --rm -v "$(pwd)"/src:/go/src -it go-4-life
```

in depth:
 - the --rm delete the container once it finished the execution
 - the -v mount the folder with the code
 - the -it 

## Folder structure

The folder structure is done checking this unofficial layout:

https://github.com/golang-standards/project-layout

## Run Tests

To run tests run the CLi command

```sh
go test -v <package>
go test -v ./internal # Example of testing the internal package
```