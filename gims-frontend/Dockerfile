FROM node:18-alpine as builder
RUN apk update ; apk add --no-cache \
  && rm -rf /var/cache/apk/*
WORKDIR /opt/

ENV NODE_ENV=production


COPY package.json /opt/
RUN npm install
COPY . . 
RUN npm run build

FROM nginx:1.21.0-alpine as production
ENV NODE_ENV production
WORKDIR /usr/share/nginx/html
COPY --from=builder /opt/build /usr/share/nginx/html/

COPY docker/etc/nginx /etc/nginx

STOPSIGNAL SIGTERM

EXPOSE 3000