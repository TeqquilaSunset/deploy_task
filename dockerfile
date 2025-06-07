FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./ 

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o /gin-backend-starter

EXPOSE 4000


FROM alpine:3.22 AS release-stage

WORKDIR /

COPY --from=build-stage /gin-backend-starter /gin-backend-starter
COPY .env .env

EXPOSE 4000

USER nobody:nobody

ENTRYPOINT ["/gin-backend-starter"]