DOCKER_USERNAME = tiangexiang
MAIN_SERVICE_VERSION = 0.1.0
MAIL_SERVICE_VERSION = 0.1.0


build_images:
	@echo "Building all the images..."
	docker build -f ./main-service/Dockerfile -t ${DOCKER_USERNAME}/main-service:${MAIN_SERVICE_VERSION} ./main-service
	docker build -f ./mail-service/Dockerfile -t ${DOCKER_USERNAME}/mail-service:${MAIL_SERVICE_VERSION} ./mail-service

remove_images:
	@echo "Removing all the images..."
	docker rmi ${DOCKER_USERNAME}/main-service:${MAIN_SERVICE_VERSION}
	docker rmi ${DOCKER_USERNAME}/mail-service:${MAIL_SERVICE_VERSION}

push_images:
	@echo "Pushing all the images..."
	docker push ${DOCKER_USERNAME}/main-service:${MAIN_SERVICE_VERSION}
	docker push ${DOCKER_USERNAME}/mail-service:${MAIL_SERVICE_VERSION}

.PHONY: build_images remove_images push_images
