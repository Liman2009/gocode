# build stage
FROM golang:alpine AS build
RUN apk update && apk add git
ADD . /src
WORKDIR /src
ENV CGO_ENABLED 0
RUN go build \
  -trimpath \
  -ldflags "-s -w -X main.version=$(git describe --abbrev=0 --tags)" \
  -o /tmp/bin
# run stage
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=build /tmp/bin /app/bin
ENTRYPOINT ["/app/bin"]
