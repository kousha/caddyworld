ARG APP_NAME="caddyworld"
ARG APP_LOCATION="github.com/kousha/"
ARG FQ_APP_NAME="${APP_LOCATION}${APP_NAME}"
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang AS build-env
ARG APP_NAME
ARG FQ_APP_NAME

RUN mkdir /workdir
WORKDIR /workdir

#Copy over go mods so that we can cache module downloads
COPY go.mod .
COPY go.sum .
RUN go mod download -x
# Copy the rest of the files
COPY . .

RUN CGO_ENABLED=0 go build

FROM scratch
ARG APP_NAME
ARG PORT=2019

COPY --from=build-env /workdir/${APP_NAME} .
COPY --from=build-env /workdir/config.json .

# Run the outyet command by default when the container starts.
CMD ["/caddyworld", "run", "--config", "config.json"]

# Document that the service listens on port 8080.
EXPOSE ${PORT}
