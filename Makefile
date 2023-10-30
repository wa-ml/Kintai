build:
	docker compose build --no-cache

up:
	docker compose up -d

front:
	docker compose up -d angular_app

backend:
	docker compose up -d go_app
down:
	docker compose down
