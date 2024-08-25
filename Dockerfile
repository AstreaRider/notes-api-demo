FROM golang:1.22-alpine3.20 AS build
WORKDIR /notes-api-demo
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server main.go

FROM alpine:3.20
WORKDIR /notes-api-demo
COPY --from=build /notes-api-demo/server .
EXPOSE 3000
CMD [ "./server" ]