# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

# Define and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy and compile source
COPY *.go ./
COPY accounts/ ./accounts/
COPY app/ ./app/
COPY cmd/ ./cmd/
COPY dccs/ ./dccs/
COPY ynab ./ynab/

RUN CGO_ENABLED=0 GOOS=linux go build -o /dccs-to-ynab

# Deploy the app binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /dccs-to-ynab ./dccs-to-ynab

# Run the sync command on an interval
ENTRYPOINT ["/dccs-to-ynab" ]
