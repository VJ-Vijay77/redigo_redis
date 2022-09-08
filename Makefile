up:
	echo "starting container"
	docker compose down
	docker compose up -d

down:
	echo "stopping containers"
	docker compose down