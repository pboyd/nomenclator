FROM golang:1.19-bullseye as builder

WORKDir /app/src

COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY . .
RUN go build -o /go/bin/api .

FROM gcr.io/distroless/base-debian11

EXPOSE 8080

COPY internal/database/migrations /app/migrations

COPY --from=builder /go/bin/api /app/bin/api

ENTRYPOINT ["/app/bin/api"]
