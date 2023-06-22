package database

//import _"path"

var connection string

func init() {
	connection = "Mongodb"
}

func getDatabse() string {
	return connection
}

// go bulid hello-world.go
// go run hello-world.go
