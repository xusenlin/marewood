# Build frontend dist.
FROM node:14-alpine AS frontend
WORKDIR /frontend-build

COPY ./public .


RUN yarn && yarn build

# Build backend exec file.
FROM golang:1.21-alpine AS backend
WORKDIR /backend-build

RUN apk add build-base

COPY . .

RUN CGO_ENABLED=1 go build -o MareWood ./MareWood.go




FROM node:16-alpine AS marewood
WORKDIR /marewood

RUN apk add git


COPY --from=backend /backend-build/MareWood /marewood
COPY --from=frontend /frontend-build/build /marewood/public/build


EXPOSE 8088
VOLUME  /marewood/resources

ENTRYPOINT ["/marewood/MareWood"]

#docker run -d --name marewood -p 8088:8088 -v ~/.marewood:/marewood/resources ghcr.io/xusenlin/marewood:0.4