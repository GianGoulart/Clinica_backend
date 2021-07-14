FROM gcr.io/distroless/base-debian10

WORKDIR /usr/src/app

COPY server .

# Run executable
CMD ["/usr/src/app/server"]