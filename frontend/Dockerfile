FROM node:14-alpine as builder
WORKDIR /app
ADD . .
RUN npm install
RUN npm run build-k8s

FROM nginx:1.20.2
COPY --from=builder /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
