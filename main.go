package main

func main() {
	server := ApiServer{":3001"}
	server.Run()
}
