.PHONY: depend, build, run, clean, test, coverageforlinux

OUT_FILENAME=exe
GOCMD=go

depend:
	go get -v -t -d ./...
	# if [ -f Gopkg.toml ]; then
    #     curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    #     dep ensure
    # fi

build:
	$(GOCMD) build -v -o $(OUT_FILENAME)

run:
	$(GOCMD) build -v -o $(OUT_FILENAME)
	./$(OUT_FILENAME)

clean:
	rm ./$(OUT_FILENAME)

test:
	$(GOCMD) test -v ./...

coverageforlinux:
	set -e \
        echo "" > coverage.txt \
        for d in $(go list ./... | grep -v vendor); do \
            go test -race -coverprofile=profile.out -covermode=atomic "$d" \
            if [ -f profile.out ]; then \
                cat profile.out >> coverage.txt \
                rm profile.out \
            fi \
        done