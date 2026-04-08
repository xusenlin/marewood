ARG NODE_VERSION=22

# Build frontend dist.
FROM node:18-alpine AS frontend
WORKDIR /frontend-build

COPY ./ui .

RUN yarn config set registry https://registry.npmmirror.com && \
    yarn && yarn build

# Build backend exec file.
FROM golang:1.21-alpine AS backend
WORKDIR /backend-build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -o main ./marewood.go


FROM node:${NODE_VERSION}-alpine AS marewood
RUN addgroup -g 1001 -S marewood && \
    adduser -u 1001 -S marewood -G marewood

WORKDIR /marewood

RUN apk add --no-cache git && \
    npm config set registry https://registry.npmmirror.com && \
    yarn config set registry https://registry.npmmirror.com && \
    if node -v | grep -q '^v16'; then \
        npm install -g pnpm@7; \
    else \
        npm install -g pnpm; \
    fi && \
    pnpm config set registry https://registry.npmmirror.com

COPY --from=backend /backend-build/main /marewood/main
COPY --from=frontend /frontend-build/dist /marewood/ui/dist

RUN mkdir -p /marewood/resources && \
    chown -R marewood:marewood /marewood

USER marewood

EXPOSE 8088
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8088/health || exit 1

VOLUME ["/marewood/resources"]

ENTRYPOINT ["/marewood/main"]