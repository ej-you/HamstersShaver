# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN go build -o ./sse_api_app ./main.go

# ---
# RUN
# ---

FROM alpine:3.20 AS run

# set up timezone +3
RUN apk update
RUN apk add --no-cache tzdata
RUN ln -s /usr/share/zoneinfo/Etc/GMT-3 /etc/localtime

# make dir for logs
RUN mkdir /logs

WORKDIR /app
# copy compiled file, files for swagger and wallet data
COPY --from=build /go/src/sse_api_app .

# run app
CMD ["/bin/sh", "-c", "/app/sse_api_app 2>> /logs/error-log.log"]
