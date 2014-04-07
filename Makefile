
NAME := "github.com/kare/gocat"

build:
	go build $(NAME)

test: build
	go test -v $(NAME)

bench: build
	go test -v -bench=. $(NAME)

clean:
	rm -rf gocat
