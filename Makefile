
NAME := "github.com/kare/gocat"

build:
	go build $(NAME)

test: build
	go test -v $(NAME)

clean:
	rm -rf gocat
