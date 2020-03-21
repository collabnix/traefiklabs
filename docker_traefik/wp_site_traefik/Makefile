release:
	rm -rf ./artifacts
	mkdir ./artifacts
	GOOS=linux GOARCH=amd64 go build -o ./artifacts/pekka-linux-amd64
	GOOS=linux GOARCH=386 go build -o ./artifacts/pekka-linux-386
	GOOS=darwin GOARCH=amd64 go build -o ./artifacts/pekka-darwin-amd64
	GOOS=darwin GOARCH=386 go build -o ./artifacts/pekka-darwin-386
	# GOOS=windows GOARCH=amd64 go build -o ./artifacts/pekka-windows-amd64.exe
	# GOOS=windows GOARCH=386 go build -o ./artifacts/pekka-windows-386.exe
