FROM golang:1.20 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine
WORKDIR /
COPY --from=build /app/cmd/api ./
EXPOSE 8080

CMD ["./api"]