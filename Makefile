build:
	docker compose build --no-cache

up:
	docker compose up -d

front:
	docker compose up -d angular_app

down:
	docker compose down