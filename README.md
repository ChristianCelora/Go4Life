# Go4Life

Conway's Game of life in golang

## Dev Env

I created a custom Docker container. 

To build it run:

```sh
docker build -t go-4-life .
```

To run it use:

```sh
docker run --rm -p 8888:8090 -v "$(pwd)"/src:/go/src -it go-4-life
```

in depth:
 - the --rm delete the container once it finished the execution
 - the -v mount the folder with the code
 - the -it 

## Folder structure

The folder structure is done checking this unofficial layout:

https://github.com/golang-standards/project-layout

## Run Tests

### Unit tests

To run tests run the CLi command

```sh
go test -v <package>
go test -v ./internal # Example of testing the internal package
```

### Integration tests

We added Jest as a framework to integration tests on APIs. 
Tests are under tests/Integration folder

To run it enter the docker CLI and run 

```sh
npm test
```

## Run web server

For now we use go to render the static html pages of the project

### Go 

Init server

```sh
go run main.go
```

Note: server should be recompiled in case of changes

## Build the FE

I added Webpack bundler to this project to import the JS modules in the app.

This has to be done because the plain HTML wasn't resolving the relative paths.

To build the FE run:

```sh
npm run build
```