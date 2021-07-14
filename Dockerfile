FROM gcr.io/distroless/base-debian10

WORKDIR /usr/src/app

EXPOSE 5055

COPY server .

# Run executable
CMD ["/usr/src/app/server"]