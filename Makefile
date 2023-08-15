build:
	docker compose up --build
auth-service:
	docker compose up --build auth_service
artist-service:
	docker compose up --build artist_service
gateway:
	docker compose up --build gateway