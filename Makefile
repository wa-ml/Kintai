build:
	docker compose build --no-cache

up:
	docker compose up -d

front:
	docker compose up angular_app

backend:
	docker compose up go_app

down:
	docker compose down
