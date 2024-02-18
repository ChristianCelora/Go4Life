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

To run tests run the CLi command

```sh
go test -v <package>
go test -v ./internal # Example of testing the internal package
```

## Run web server

For now we use go to render the static html pages of the project

### Go 

Init server

```
go run main.go
```

Note: server should be recompiled in case of changes

### Apache (not used)

Apache2 is installed in the docker. Just run it inside the container

```sh
apachectl start
```

To check the docker ip run a docker inspect

```sh
docker inspect <container_ip>
```

and check for the NetworkSettings > Ports settings for the ip and port open to connection