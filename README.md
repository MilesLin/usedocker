This is a simple docker API to call docker command on the remote machine. It includes basic authorization for calling API.


You can use `go run . -h` to see the flag can be used.

**Enable SSL**

`go run . -enableSSL`

**Enable basic authorization please set the flag to acct and pwd**

`go run . -acct=admin -pwd=123456`
 
**Specify the port**

`go run . -port=8080 -sslport=443`

the default port of http is 8080, https is 443.

**Enable swagger**

`go run . -swag`

The swagger page is http://localhost:8080/swagger/index.html. *The port depends on your command*.


