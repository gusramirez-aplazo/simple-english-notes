# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

WORKDIR /notes

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o /english-notes ./server

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /english-notes /

EXPOSE 8080

ENTRYPOINT ["/english-notes"]
