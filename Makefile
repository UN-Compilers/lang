build-docker:
	docker build --build-arg CURR_USER_UID=`id -u` --build-arg CURR_USER_GID=`id -g` -t go-antlr .
run-docker:
	docker run -it --rm -v .:/home/gousr/project go-antlr
install-dependencies:
	go mod tidy
run-tests-ci:
	go test ./own/
run-tests:
	 docker run --rm -v .:/home/gousr/project go-antlr go test ./own
generate-parsing-files:
	go generate ./...