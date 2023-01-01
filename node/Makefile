docker-build:
	DOCKER_BUILDKIT=1 docker build -t node-udp-server:latest .

docker-run:
	docker run --rm -d -p 3002:3002 -p 3002:3002/udp --name node-udp-server node-udp-server:latest
