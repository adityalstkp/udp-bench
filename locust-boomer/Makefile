DUMMY_PYTHON_FILE_PATH=/mnt/locust
LOCUST_MASTER_DOCKER_NAME=locust-boomer-master

start-locust-master:
	docker run --rm -d -p 8089:8089 -p 5557:5557 \
		--name $(LOCUST_MASTER_DOCKER_NAME) \
		-v $$PWD/dummy.py:$(DUMMY_PYTHON_FILE_PATH)/dummy.py \
		locustio/locust -f $(DUMMY_PYTHON_FILE_PATH)/dummy.py --master

stop-locust-master:
	docker kill $(LOCUST_MASTER_DOCKER_NAME) || true

build-boomer:
	go build -o bin/locust-boomer cmd/udp/main.go

serve-boomer: build-boomer
	bin/locust-boomer

swarm: stop-locust-master run-locust-master serve-boomer
