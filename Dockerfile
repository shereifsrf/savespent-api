#build stage
FROM golang:alpine
WORKDIR $GOPATH/src/github.com/shereifsrf/savespent-api
COPY . .
# Download all the dependencies
RUN go get -d -v ./...
# Install the package
RUN go install -v ./...
# Run the executable
CMD ["savespent-api"]
