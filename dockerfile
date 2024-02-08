# Build Stage

FROM golang:1.21.1 AS buildstage

RUN mkdir /app

WORKDIR /app

COPY /api /app/api 
COPY /web /app/web
COPY /main.go /app/main.go
COPY /go.mod /app/go.mod
COPY /go.sum /app/go.sum

RUN go mod download

EXPOSE 8080

RUN CGO_ENABLED=1 go build -o /app/main .

# Deploy Stage

FROM gcr.io/distroless/cc-debian12

WORKDIR /app

COPY --from=buildstage /app /app

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]