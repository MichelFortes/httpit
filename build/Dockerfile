FROM golang:1.22.5 as build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o httpit ./cmd/httpit.go

FROM alpine
WORKDIR /app
RUN apk update
RUN apk upgrade -U --no-interactive
RUN apk add curl --no-interactive
COPY --from=build /app/httpit .
ENTRYPOINT [ "./httpit" ]
