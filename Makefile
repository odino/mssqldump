DOCKER_IMAGE_NAME=mssqldump

test:
	docker run -ti \
		--env-file=env.secret \
		-e LOG_LEVEL=DEBUG \
		-v $$(pwd):/go/src/github.com/odino/mssqldump \
		--net=host \
		mssqldump \
		bash
build:
	docker build -t mssqldump .
release_simple:
	rm -rf builds/
	go build -o builds/simple
release: release_simple
	GOOS=linux GOARCH=amd64 go build -o builds/mssqldump_linux_amd64_$$(./builds/simple -qv)
	GOOS=linux GOARCH=386 go build -o builds/mssqldump_linux_386_$$(./builds/simple -qv)
	GOOS=darwin GOARCH=amd64 go build -o builds/mssqldump_darwin_amd64_$$(./builds/simple -qv)
	GOOS=darwin GOARCH=386 go build -o builds/mssqldump_darwin_386_$$(./builds/simple -qv)
	GOOS=windows GOARCH=amd64 go build -o builds/mssqldump_windows_amd64_$$(./builds/simple -qv)
	GOOS=windows GOARCH=386 go build -o builds/mssqldump_windows_386_$$(./builds/simple -qv)
	sudo chown $$USER:$$USER builds
	sudo chown $$USER:$$USER builds/*
	cd builds; ls -la . | grep -v ".tar.gz" | grep mssqldump | awk '{print "tar -czf " $$9 ".tar.gz " $$9}' | bash
	cd builds; ls -la . | grep -v ".tar.gz" | grep mssqldump | awk '{print "rm " $$9}' | bash
	ls -la builds
