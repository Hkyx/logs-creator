# Create a minimal container to run a Golang static binary

FROM alpine:latest

EXPOSE 8080

WORKDIR /app

# copy binary into image
COPY golanglog /app/

ENTRYPOINT ["./golanglog"]
