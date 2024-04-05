test-swiftmac:
	go build main.go
	./main -p ~/projects/robertmeta/swiftmac/.build/debug/swiftmac -s swiftmac.etts

test-sharpwin:
	go build main.go
	./main.exe -p ..\SharpWin\bin\Debug\net8.0\SharpWin.exe -s .\SharpWin.etts
