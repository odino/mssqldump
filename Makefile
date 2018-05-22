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
	docker run -ti --net host -v $$(pwd):/go/src/github.com/odino/mssqldump ${DOCKER_IMAGE_NAME} go build -o builds/simple main.go
release: release_simple
	docker run -ti --net host -v $$(pwd):/go/src/github.com/odino/mssqldump ${DOCKER_IMAGE_NAME} gox -output="builds/{{.Dir}}_{{.OS}}_{{.Arch}}_$$(./builds/simple -qv)"
	sudo chown $$USER:$$USER builds
	sudo chown $$USER:$$USER builds/*
	cd builds; ls -la . | grep -v ".tar.gz" | grep mssqldump | awk '{print "tar -czf " $$9 ".tar.gz " $$9}' | bash
	cd builds; ls -la . | grep -v ".tar.gz" | grep mssqldump | awk '{print "rm " $$9}' | bash
	ls -la builds
