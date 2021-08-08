FROM golang AS build

WORKDIR grpc_test_task/service_get

COPY get_interface_dumps .

RUN CGO_ENABLED=0 go build -o /bin/server
FROM scratch
COPY --from=build /bin/server /bin/server


ENTRYPOINT ["/bin/server"]
EXPOSE 50052