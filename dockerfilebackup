FROM golang:1.17.7-bullseye
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod vendor
ENV LISTEN_PORT=5000
EXPOSE 5000
RUN go build -o main .
CMD ["/app/main"]