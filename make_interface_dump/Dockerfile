FROM golang AS build

WORKDIR grpc_test_task/service_make

COPY make_interface_dump .

RUN CGO_ENABLED=0 go build -o /bin/server
FROM scratch
COPY --from=build /bin/server /bin/server


ENTRYPOINT ["/bin/server"]
EXPOSE 50051