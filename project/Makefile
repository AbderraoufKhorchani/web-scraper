SCRAPER_BINARY=scrapingApp


up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_scraper:
	@echo "Building scraper binary..."
	cd ../scraping-service && env GOOS=linux CGO_ENABLED=0 go build -o ${SCRAPER_BINARY} .
	@echo "Done!"	

up_build: build_scraper
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"	

install:
	@echo "Installing scraper dependencies..."
	cd ../scraping-service && go install ./...
