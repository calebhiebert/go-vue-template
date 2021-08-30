# Start with a golang image
FROM golang:1.17-stretch as build

ENV GO111MODULE on

# Create a user to run the app as
RUN useradd --shell /bin/bash go-vue-template

# Set the workdir to the application path
WORKDIR $GOPATH/src

# Copy all application files
COPY . ./

# Build the app
RUN cd $GOPATH/src/go-vue-template && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -a -installsuffix nocgo -ldflags="-w -s" -o /go/bin/go-vue-template

# Start from a scratch container for a nice and small image
FROM alpine:3.8

# Install ca-certificates for calling https endpoints
RUN apk add --no-cache ca-certificates && mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Copy the binary build
COPY --from=build /go/bin/go-vue-template /go/bin/go-vue-template

# Copy the password file (with the go-vue-template user) from the build container
COPY --from=build /etc/passwd /etc/passwd

# Copy the timezone information in for golang
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /go/bin

# Set time info environment variable
ENV ZONEINFO=/go/bin/zoneinfo.zip

# Set the user to the previously created user
USER go-vue-template

# Set the workdir
WORKDIR /go/bin

# Expose the API port
EXPOSE 8080

CMD [ "/go/bin/go-vue-template" ]