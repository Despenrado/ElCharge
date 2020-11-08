FROM golang:1.15 AS builder
WORKDIR /go/src/github.com/Ddespenrado/ElCharge/RestAPI/
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o restapi .

FROM scratch
EXPOSE 8081
WORKDIR /root/
COPY deploy/apiserver.yaml .
COPY --from=builder /go/src/github.com/Ddespenrado/ElCharge/RestAPI/restapi .
ENTRYPOINT ["./restapi","-config=apiserver.yaml"]