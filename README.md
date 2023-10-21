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