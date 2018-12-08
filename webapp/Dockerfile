FROM golang:1.11-stretch

# Install OS libs
RUN apt-get update -y && \
  apt-get install -y \
    curl \
    git \
    wget \
    xmlsec1
# Install Go's dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

###################
# App Preparation #
###################
# Set go working dir for this app
WORKDIR /go/src/github.com/aren55555/learn/dec7/webapp
# Copy dep files
COPY Gopkg.lock .
COPY Gopkg.toml .
RUN dep ensure \
    -v \
    -vendor-only

###################
# App Preparation #
###################
# Copy all of the app
COPY . .
# Compile binary, and put in PATH
RUN go install -v ./...

# Exec go bin
CMD ["webapp"]
