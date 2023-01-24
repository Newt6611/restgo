FROM golang:1.19-alpine
# WORKDIR /server
# COPY . .
RUN go mod download

