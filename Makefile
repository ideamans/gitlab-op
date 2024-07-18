gitlab-op: *.go go.mod go.sum
	go build -o gitlab-op ./...

.PHONY: install
install: gitlab-op
	sudo cp gitlab-op /usr/local/bin