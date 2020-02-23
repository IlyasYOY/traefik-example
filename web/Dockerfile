FROM node:12-alpine as builder

WORKDIR /opt/build

COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json
COPY ./.babelrc ./.babelrc
COPY ./webpack.config.js ./webpack.config.js
RUN npm install

COPY ./src ./src
RUN npm run build

FROM nginx:alpine

COPY --from=builder /opt/build/dist /usr/share/nginx/html

COPY  ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]

