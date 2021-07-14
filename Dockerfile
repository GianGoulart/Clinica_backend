FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY server .
COPY ./config.json .

# Run executable
CMD ["/app/server"]