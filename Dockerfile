FROM golang:1.21.4-alpine
WORKDIR /otpapi
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/main .
CMD ./out/main
