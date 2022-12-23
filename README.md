# Simple Socket Chat

Terminal chat application 

## Usage

This application implements both server and client.

### Server

To start the server enter 
```shell
go run main.go -service server
```
You can also specify the port to run using the key `-port` (default `8080`)

### Client
To start the client enter
```shell
go run main.go
```
Similar to the server, you can specify the server port to connect to (default `8080`)

## Example

An example of using the application and its output is given

### Server
As you can see, the server logs all messages and writes to the chat, starting with `>`
```text
2022/12/24 10:13:55 server started on :8080
2022/12/24 10:14:18 > [Gosha] connected, 127.0.0.1:58861
2022/12/24 10:14:45 > [Fill] connected, 127.0.0.1:58876
2022/12/24 10:15:00 [Gosha] Hi, Fill!, 127.0.0.1:58861
2022/12/24 10:15:20 [Fill] Hello, 127.0.0.1:58876
2022/12/24 10:16:05 [Gosha] I'm from Russia, 127.0.0.1:58861
2022/12/24 10:16:12 > [Fill] left, 127.0.0.1:58876
```
### Client
Immediately after starting the client, you need to enter a nickname. Then server add new client to the chat. System messages also start with `>`

```text
> Enter your nickname: Gosha
> To exit enter Ctrl+C
> [Fill] connected
Hi, Fill!
[Fill] Hello
I'm from Russia 
> [Fill] left
```
