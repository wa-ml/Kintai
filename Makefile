build:
	docker compose build --no-cache

up:
	docker compose up -d

front:
	docker compose up angular_app

back:
	docker compose up go_app mysql phpmyadmin mailhog

go:
	docker compose up go_app

down:
	docker compose down
