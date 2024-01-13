build:
	go build -o bin/myapp
run: build
	./bin/myapp
git:
	git init