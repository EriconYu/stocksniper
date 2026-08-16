.PHONY: install dev test lint security check docker-build docker-run

install:
	python3 -m pip install -r requirements.txt

dev:
	STOCKSNIPER_RELOAD=true python3 -m app.main

test:
	python3 -m pytest

lint:
	python3 -m ruff check .

security:
	python3 scripts/security_check.py

check: lint test security

docker-build:
	docker build -t stocksniper-community:latest .

docker-run:
	docker run --rm -p 8000:8000 stocksniper-community:latest
