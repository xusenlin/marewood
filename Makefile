# ========== 使用说明 ==========
# make pull          # 拉取基础镜像
# make build         # 构建所有版本镜像 (node16-24)
# make build-node24  # 构建 node24 镜像
# make check         # 运行最新版本镜像测试
# make check-node24  # 运行 node24 镜像测试
# make login TOKEN=your_token   # 登录 ghcr.io
# make push          # 推送所有版本镜像
# make push-node24   # 推送 node24 镜像
# ========== 配置 ==========
BASE_IMAGES := node:16-alpine node:18-alpine node:20-alpine node:22-alpine node:24-alpine golang:1.21-alpine
DOCKER_USERNAME := xusenlin
PROJECT_NAME := marewood
PROJECT_VERSION := 1.2
BUILD_NODE_IMAGE := 16  18  20  22  24

# 获取项目全部基础镜像
pull:
	@for image in $(BASE_IMAGES); do \
		echo "\nPulling base image for $$image..."; \
		docker pull $$image; \
	done

build:
	@for nodeVersion in $(BUILD_NODE_IMAGE); do \
        imageName=ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node$$nodeVersion; \
		echo "\nBuilding Docker image for $$imageName..."; \
		docker build --build-arg NODE_VERSION=$$nodeVersion -t $$imageName .; \
    done

build-node24:
	@imageName=ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node24; \
	echo "\nBuilding Docker image for $$imageName..."; \
	docker build --build-arg NODE_VERSION=24 -t $$imageName .

check:
	docker run -d --name marewood -p 8088:8088 -v marewood-data-$(lastword $(BUILD_NODE_IMAGE)):/marewood/resources ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node$(lastword $(BUILD_NODE_IMAGE))

check-node24:
	docker run -d --name marewood -p 8088:8088 -v marewood-data-24:/marewood/resources ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node24

push:
	@for nodeVersion in $(BUILD_NODE_IMAGE); do \
        imageName=ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node$$nodeVersion; \
		echo "\nPushing Docker image $$imageName..."; \
		docker push $$imageName; \
    done

push-node24:
	@imageName=ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node24; \
	echo "\nPushing Docker image $$imageName..."; \
	docker push $$imageName

login:
	@echo "$(TOKEN)" | docker login ghcr.io -u $(DOCKER_USERNAME) --password-stdin


.PHONY: pull build build-node24 check check-node24 push push-node24 login