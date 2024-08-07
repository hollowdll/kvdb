# API

The kvdb server API is implemented with gRPC and defined with Protocol Buffers. HTTP/2 is needed in the connections. Requests are made with RPCs (Remote Procedure Calls). Connections require a gRPC client.

RPCs are grouped to gRPC services that provide related RPCs. The protobuf gRPC service definitions are in the `api/` directory. It contains the currently maintained API version. All services and RPCs are documented there in files that have `.proto` file extension.

[Link to the API directory](../api/)

# Authentication

If the server is password protected, the client needs to authenticate with password. The password is sent in gRPC metadata and needs to be included in all requests.

The gRPC metadata key for password is `password`. The actual password to send is set as the value of the key.

# API v0 gRPC services

## ServerService

- Service name: `ServerService`
- Package: `api.v0.serverpb`
- Proto file: `server.proto`

[Link to the protobuf definitions](../api/v0/serverpb/server.proto)

ServerService provides server-related RPCs.

Common gRPC metadata for this service's RPCs:
- `password`: The server password. Used for authentication if the server is password protected.

## DBService

- Service name: `DBService`
- Package: `api.v0.dbpb`
- Proto file: `db.proto`

[Link to the protobuf definitions](../api/v0/dbpb/db.proto)

DBService provides database-related RPCs.

Common gRPC metadata for this service's RPCs:
- `password`: The server password. Used for authentication if the server is password protected.

## GeneralKVService

- Service name: `GeneralKVService`
- Package: `api.v0.kvpb`
- Proto file: `general_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/general_kv.proto)

GeneralKVService provides RPCs related to general key management.

Common gRPC metadata keys for this service's RPCs:
- `password`: The server password. Used for authentication if the server is password protected.
- `database`: The database to use. If omitted, the default database is used.

## StringKVService

- Service name: `StringKVService`
- Package: `api.v0.kvpb`
- Proto file: `string_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/string_kv.proto)

StringKVService provides RPCs related to String keys.

Common gRPC metadata keys for this service's RPCs:
- `password`: The server password. Used for authentication if the server is password protected.
- `database`: The database to use. If omitted, the default database is used.

## HashMapKVService

- Service name: `HashMapKVService`
- Package: `api.v0.kvpb`
- Proto file: `hashmap_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/hashmap_kv.proto)

HashMapKVService provides RPCs related to HashMap keys.

Common gRPC metadata keys for this service's RPCs:
- `password`: The server password. Used for authentication if the server is password protected.
- `database`: The database to use. If omitted, the default database is used.

