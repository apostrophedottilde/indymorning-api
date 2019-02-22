# FROM golang:1.11-alpine AS builder
# RUN apk update && apk add --no-cache && apk add gcc && apk add musl-dev && apk add git && mkdir app && apk add tree
# COPY . /app
# WORKDIR /app
# RUN go get -d -v
# RUN go mod download
# RUN go build -o /go/bin/projects
# RUN ls


# FROM scratch
# COPY --from=builder /go/bin/projects /go/bin/projects
# ENTRYPOINT ["/go/bin/projects"]

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache && apk add gcc && apk add musl-dev && apk add git && mkdir app
COPY . /app
WORKDIR /app
RUN go get -d -v
RUN go mod download
RUN go build
ENTRYPOINT ["go", "run", "main.go"]