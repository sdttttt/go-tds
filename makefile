OUT_FILENAME=exe
GOCMD=go

build:
	$(GOCMD) build -o $(OUT_FILENAME)

run:
	$(GOCMD) build -o $(OUT_FILENAME)
	./$(OUT_FILENAME)

clean:
	rm ./$(OUT_FILENAME)

test:
	$(GOCMD) test ./...