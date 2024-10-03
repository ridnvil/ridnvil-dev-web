# Use a minimal base image for Go applications
FROM golang:1.22.7-alpine AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .
COPY ridnvil /app/ridnvil
COPY database /app/database
COPY nvil.sqlite3 /app/nvil.sqlite3
COPY ridnvil/build/static/css /app/static/css
COPY ridnvil/build/static/js /app/static/js
COPY ridnvil/build/static/media /app/static/media

RUN apk --no-cache add tzdata
ENV TZ=Asia/Jakarta
ENV CGO_ENABLED=1

EXPOSE 3001
CMD ["./main"]