FROM node:lts-alpine as FPack-base
RUN apk add --no-cache git build-base
WORKDIR /marewood-f

FROM FPack-base as FBuild-vendor-pre
COPY public/package.json /marewood-f/package.json
RUN node -pe '(d=`dependencies`,dd=`devDependencies`,p=`./package.json`,JSON.stringify({[d]:require(p)[d],[dd]:require(p)[dd]},null,2))' >pkg-build-deps.json

FROM FPack-base as FBuild-vendor-build
COPY public/yarn.lock /marewood-f/yarn.lock
COPY --from=FBuild-vendor-pre /marewood-f/pkg-build-deps.json /marewood-f/package.json
RUN yarn install

FROM FPack-base as FBuild
COPY public /marewood-f
COPY --from=FBuild-vendor-build /marewood-f/node_modules /marewood-f/node_modules
ENV REACT_APP_API_URL=/
RUN yarn build

FROM golang:1.13-alpine as Build-pre
WORKDIR /marewood
COPY . /marewood
RUN  rm -rf public docker-run.sh

FROM golang:1.13-alpine as Build
RUN apk add --no-cache git build-base
WORKDIR /marewood
COPY --from=Build-pre /marewood /marewood
RUN  go build -mod=vendor -o marewood

FROM node:lts-alpine
RUN apk add --no-cache git
WORKDIR /marewood/data
COPY --from=Build /marewood/marewood /marewood/marewood
COPY --from=FBuild /marewood-f/build /marewood/public/build
COPY docker-run.sh /marewood/docker-run.sh
RUN chmod +x /marewood/docker-run.sh
ENV LogMode=true \
  GinMode=debug \
  HttpPort=8081 \
  AddressUrl=http://localhost:8081 \
  ClientDir=/marewood/public/build
CMD [ "/marewood/docker-run.sh" ]
