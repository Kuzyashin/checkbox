FROM golang:alpine AS builder
RUN apk update && apk add --no-cache ca-certificates git gcc make libc-dev binutils-gold
RUN mkdir Checkbox
WORKDIR Checkbox
COPY ../cmd cmd
COPY ../internal internal
COPY ../pkg pkg
COPY ../go.mod go.mod
COPY ../go.sum go.sum
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g cmd/main/main.go --output docs/
WORKDIR cmd/main
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /bin/Checkbox

FROM scratch
COPY --from=builder /bin/Checkbox /bin/Checkbox
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/bin/Checkbox"]