FROM golang:1.20-bookworm

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY go.mod go.sum ./
COPY . ./

# Download Go modules
RUN go mod tidy
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /psi-anamnese-be

# Run
CMD ["/psi-anamnese-be"]
