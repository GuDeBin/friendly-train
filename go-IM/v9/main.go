package main

func main() {
	// v1版本
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
