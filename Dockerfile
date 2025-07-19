FROM    golang:1.24-alpine AS dependencies
WORKDIR /opt/app-code
COPY    ./go.mod .
COPY    ./go.sum .
COPY    ./.env.app .
RUN     go mod download

FROM    golang:1.24-alpine AS builder
COPY    ./ /opt/app-code/
COPY    --from=dependencies /opt/app-code/ /opt/app-code/
WORKDIR /opt/app-code/
RUN     go build -v -o ./dist/app ./main.go

FROM    alpine:3.20 AS dist
COPY    --from=builder /opt/app-code/dist/. /opt/app-code
COPY    --from=builder /opt/app-code/.env.app /opt/app-code
WORKDIR /opt/app-code
CMD     ["./app"]