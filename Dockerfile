FROM golang:1.20-buster AS build
RUN go install github.com/msw-x/vgen/cmd/vgen@latest
WORKDIR /app
COPY . ./
RUN vgen go src
WORKDIR /app/src
RUN go mod download
RUN go build -o /sws

FROM gcr.io/distroless/base-debian10
COPY --from=build /sws ./
COPY routes.lst ./etc/sws/routes.lst
COPY sws.conf ./
ENTRYPOINT ["/sws"]
