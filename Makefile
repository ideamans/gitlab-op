gitlab-op:
	go build -o gitlab-op ./...

.PHONY: install
install: gitlab-op
	sudo cp gitlab-op /usr/local/bin