
NAME := "github.com/kare/gocat"

build:
	go build $(NAME)

test:
	go test -v $(NAME)

clean:
	rm -rf gocat
