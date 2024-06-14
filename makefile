run-frontend:
	cd something && npm run dev 

run-backend:
	templ generate && go run .

run: run-frontend run-backend

build:
	docker build buildx -t dserrano/solpulse:0.0.0 . && docker run