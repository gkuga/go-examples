FROM golang:1.23.2 AS build

WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o counter

FROM alpine:3.15
WORKDIR /app
COPY --from=build /app/counter .
CMD ["./counter"]
