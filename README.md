# Go Webhook Logger

This is a simple webhook logger written in Go. It listens for incoming HTTP requests and logs them to a file in markdown.

## Usage

Run the server with the following command and choose a port to listen on.

```bash
$ go run main.go --port=8080
```
  
Then, send a POST request to the server with the following JSON body.

example url: http://localhost:8080/listen

```json
{
  "message": "Hello, world!"
}
```

The server will log the request to a file in the `logs` directory. The file will be named with the current timestamp.
