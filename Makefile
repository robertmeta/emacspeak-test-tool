test-swiftmac:
	go build main.go
	./main -p ~/projects/robertmeta/swiftmac/.build/debug/swiftmac -s swiftmac.etts
