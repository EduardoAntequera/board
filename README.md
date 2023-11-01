# board
Testing out an example client where a cli board gets a position from a local server

# How to run
You should be running locally the board-server
Go to folder and from terminal `go run .`
Or
Create a go build and use that instead to run the program

You can use `watch -n 1 go run .` to see the board update on the cli every second since it will request a new position each time.
