# Stage 1: Build Go application
FROM golang:1.23.4-alpine AS build-go

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Build ReactJS app
FROM node:20-alpine AS build-react

WORKDIR /ridnvil

# Salin hanya file yang dibutuhkan untuk install dan build
COPY ridnvil/package*.json ./

RUN npm install

# Salin seluruh source code React
COPY ridnvil ./

RUN npm run build

# Stage 3: Final image
FROM alpine:latest

WORKDIR /app

# Copy Go binary
COPY --from=build-go /app/main .

# Copy supporting files
COPY ridnvil /app/ridnvil
COPY database /app/database
COPY nvil.sqlite3 /app/nvil.sqlite3

# Copy React build output to desired location (e.g., serve via Go or static server)
COPY --from=build-react /ridnvil/build /app/ridnvil/build

# Set timezone and environment
RUN apk --no-cache add tzdata
ENV TZ=Asia/Jakarta
ENV CGO_ENABLED=1

EXPOSE 3001
CMD ["./main"]
