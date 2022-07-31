FROM golang:bullseye

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary .

RUN chmod a+x binary .

ENTRYPOINT ["/app/binary"]
