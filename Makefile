test: actual.txt
	diff -u expected.txt actual.txt

actual.txt: streamer inp.json
	cat inp.json | ./streamer > actual.txt

streamer: main.go
	go build -o streamer
