TAGS=v1.1
DOCKERFILE_DIR=.
PREFIX=github.com/EriconYu
all: build 
build:
	docker build  -t stocksniper:$(TAGS)  .
.PHONY: build 
