FROM golang


ENV GOPATH=/

COPY ./ ./


RUN go mod download
RUN go build -o cat-app ./cmd/main.go

CMD ["./cat-app"]