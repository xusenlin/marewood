# Build frontend dist.
FROM node:18-alpine AS frontend
WORKDIR /frontend-build

COPY ./ui .

RUN yarn config set registry https://registry.npmmirror.com
RUN yarn && yarn build

# Build backend exec file.
FROM golang:1.21-alpine AS backend
WORKDIR /backend-build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main ./marewood.go


FROM node:22-alpine AS marewood
WORKDIR /marewood

#(node > 18)
RUN corepack enable && corepack prepare pnpm@latest --activate

RUN npm config set registry https://registry.npmmirror.com
RUN yarn config set registry https://registry.npmmirror.com
RUN pnpm config set registry https://registry.npmmirror.com


RUN apk add git


COPY --from=backend /backend-build/main /marewood/main
COPY --from=frontend /frontend-build/dist /marewood/ui/dist


EXPOSE 8088
VOLUME  /marewood/resources

ENTRYPOINT ["/marewood/main"]