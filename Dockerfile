FROM golang:1.21.0 as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /rinha cmd/main.go

FROM scratch
COPY --from=builder /rinha /rinha

EXPOSE 8181
ENTRYPOINT [ "/rinha" ]