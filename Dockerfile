FROM golang:1.16-buster AS build
WORKDIR /app
COPY src/* ./
RUN go mod download
RUN go build -o /sws

FROM gcr.io/distroless/base-debian10
COPY --from=build /sws ./
COPY routes.lst ./etc/sws/routes.lst
COPY sws.conf ./
ENTRYPOINT ["/sws"]
