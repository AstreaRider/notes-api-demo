FROM golang:1.22-alpine3.20 AS build
WORKDIR /notes-api-demo
COPY go.mod go.sum ./
RUN go mod download
COPY config/ ./config
COPY database/ ./database
COPY internal/ ./internal
COPY router/ ./router
COPY main.go ./
RUN go build -o server main.go

FROM alpine:3.20
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
WORKDIR /notes-api-demo
COPY --from=build /notes-api-demo/server .
EXPOSE 3000
USER nonroot
CMD [ "./server" ]