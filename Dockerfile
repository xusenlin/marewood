# Build frontend dist.
FROM node:18-alpine AS frontend
WORKDIR /frontend-build

COPY ./ui .

RUN yarn && yarn build

# Build backend exec file.
FROM golang:1.21-alpine AS backend
WORKDIR /backend-build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main ./marewood.go


FROM node:20-alpine AS marewood
WORKDIR /marewood

#(node16)
#RUN npm install -g pnpm@7 && npm cache clean --force


RUN corepack enable && corepack prepare pnpm@latest --activate
RUN apk add git

COPY --from=backend /backend-build/main /marewood/main
COPY --from=frontend /frontend-build/dist /marewood/ui/dist


EXPOSE 8088
VOLUME  /marewood/resources

ENTRYPOINT ["/marewood/main"]

#docker run -d --name marewood -p 8088:8088 -v ~/docker/marewood:/marewood/resources ghcr.io/xusenlin/marewood:1.0.0-node18