.PHONY: coverage install
coverage:
	go test -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
	echo "you should checkout coverage.html in your browser"
clean:
	rm c.out
	rm coverage.html
install:
	go install github.com/shanemhansen/ballclock/...
