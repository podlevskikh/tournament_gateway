PROJECT?=volleymsk
APP?=tournament_gateway
PORT?=80
PORT_APP?=7784

CONTAINER_IMAGE?=$(PROJECT)/${APP}
RELEASE?=0.0.1

clean:
	rm -f bin/${APP}

gorun: clean
	go build -o bin/${APP} -tags "dev load_envs" ./cmd/ && bin/${APP}

container:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && docker rm $(CONTAINER_IMAGE):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT_APP} --rm \
		-e "PORT=${PORT}" \
		--env-file .env  \
		$(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -tags="testing" -v -race -cover -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out
