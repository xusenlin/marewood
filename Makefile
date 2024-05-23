# 定义可用的 Node 版本
BASE_IMAGES := node:16-alpine node:18-alpine node:20-alpine node:22-alpine golang:1.21-alpine
# Docker Hub 用户名
DOCKER_USERNAME := xusenlin
# 项目名称
PROJECT_NAME := marewood
# 项目版本
PROJECT_VERSION := 1.0.3
# 构建的镜像
#node 16 因为没有pnpm,所以需要手动构建
BUILD_NODE_IMAGE := 18  20  22

# 获取项目全部基础镜像
pull:
	@for image in $(BASE_IMAGES); do \
		echo "\nPulling base image for $$image..."; \
		docker pull $$image; \
	done

# mac 和 linux的 sed 命令有差异，需要注意,目前在mac下测试ok.
build:
	@for nodeVersion in $(BUILD_NODE_IMAGE); do \
        imageName=ghcr.io/$(DOCKER_USERNAME)/$(PROJECT_NAME):$(PROJECT_VERSION)-node$$nodeVersion; \
		echo "\nBuilding Docker image for $$imageName..."; \
		sed -i "" "s/FROM node:[0-9][0-9]-alpine AS marewood/FROM node:$$nodeVersion-alpine AS marewood/g" Dockerfile; \
		sudo docker build -t $$imageName .; \
    done


check:
	sudo docker run -d --name marewood -p 8088:8088 -v ~/docker/marewood:/marewood/resources ghcr.io/xusenlin/marewood:$(PROJECT_VERSION)-node22


.PHONY: pull build check