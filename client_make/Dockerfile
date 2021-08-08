FROM golang AS build

WORKDIR grpc_test_task/client_make

COPY client_make .

RUN CGO_ENABLED=0 go build -o /bin/client
FROM scratch
COPY --from=build /bin/client /bin/client


ENTRYPOINT ["/bin/client"]
EXPOSE 50051