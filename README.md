# mcserver-backend
The backend for a Minecraft Server hosting platform.

## Usage
An API server listens on port 8000 by default. use the `./main --api-port new-port` to run the server on a different port.
The server has the following routes:

**createServer** - create a new Minecraft server.

| Attribute | Value  |
| --------- | ------ |
| Method    | `POST` |
| Endpoint  | `/api/createServer/` |

Request:
```
{
  "version": "1.12",
  "name": "Awesome server name",
  "port": "25565",
  "ram": "1000"
}
```

**changeProperty** - change a value in the `server.properties` file of a server.

| Attribute | Value  |
| --------- | ------ |
| Method    | `POST` |
| Endpoint  | `/api/changeProperty/` |

Request:
```
{
  "hash": "5daa124f",
  "property": "ServerPort",
  "newValue": "25599"
}
```

**system** - execute a systemctl command (either `start`, `stop`, or `status`)

| Attribute | Value  |
| --------- | ------ |
| Method    | `GET`  |
| Endpoint  | `/api/system/{method}/{hash}` |

`{method}` is the systemctl command and `{hash}` is the target server's hash or id.

**deleteServer** - delete a server

| Attribute | Value  |
| --------- | ------ |
| Method    | `GET` |
| Endpoint  | `/api/deleteServer/{hash}` |

`{hash}` is the target server's hash or id.