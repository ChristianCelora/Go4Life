FROM ubuntu:jammy

ARG VERSION="1.21.1" # go version
ARG ARCH="arm64" # go archicture
ARG NODE_VERSION=20 # node version

RUN apt-get update -y && \
    apt install curl -y && \
    apt-get clean 

# Install node through NVM (Node Version Manager)
# Change shell to use "source" cmd (-i fixes a bug with ubuntu image)
SHELL ["/bin/bash", "--login", "-i", "-c"]
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash && \
    source /root/.bashrc && \ 
    nvm install ${NODE_VERSION} && \
    node -v && \
    npm -v
SHELL ["/bin/bash", "--login", "-c"]

# Get go binary
RUN curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"

# Extract 
RUN tar -C /usr/local -xf "go${VERSION}.linux-${ARCH}.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH

EXPOSE 80 
EXPOSE 8090