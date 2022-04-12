FROM golang:1.18-alpine as BUILD-STAGE

WORKDIR /src

COPY . .
RUN go mod download
RUN go build -o /app/backend

FROM golang:1.18-alpine as RUN-STAGE

WORKDIR /app

COPY --from=BUILD-STAGE . .

CMD ["/app/backend"]