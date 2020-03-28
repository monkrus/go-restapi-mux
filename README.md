# go-restapi-mux
Building a REST API with Go lang and the MUX router

1. Create a entry point of the app.(`server.go`). 
   Within the main function we are going to create a new instance of mux router.

2. Install Mux dependency `go get -u github.com/gorilla/mux`
   Add imports and create a new router.

3. Create handler to check that our server is up and running.
   In there we have a root path and inline function that receives response and request.

4. Create the port constant. Listen and serve using port and router.
   Create a log for "server listening on the port..."
   Test the request/response and print the results.(log.Fatalln)

5. Run `go build` and `go run .` to test the server functionality.


