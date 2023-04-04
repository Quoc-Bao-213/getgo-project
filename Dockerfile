# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20.2-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY api/go.mod .

# copy directory files i.e all files ending with .go
COPY api .

# download Go modules and dependencies
RUN go mod download

# compile application
# /getgoapi: directory stores binaries file
RUN go build -o /getgoapi ./cmd/getgoserver/main.go ./cmd/getgoserver/router.go

##
## STEP 2 - DEPLOY
##
FROM scratch
WORKDIR /
COPY --from=builder /getgoapi /getgoapi

ENTRYPOINT ["./getgoapi"]
