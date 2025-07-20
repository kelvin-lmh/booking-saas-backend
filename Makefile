build:
	docker compose -f docker-compose.yml -f docker-compose.postgres.yml up -d ${SERVICE}

down:
	docker compose -f docker-compose.yml -f docker-compose.postgres.yml down ${SERVICE}