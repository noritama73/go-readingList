FROM node:lts-alpine as builder
WORKDIR /app
COPY ./frontend/package.json ./
RUN yarn install
COPY ./frontend .
RUN yarn build
RUN mv dist /

FROM nginx:stable-alpine 
COPY --from=builder /dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]