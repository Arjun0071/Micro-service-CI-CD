# Stage 1: Base image
FROM nginx:alpine

# Stage 2: Remove default nginx static files
RUN rm -rf /usr/share/nginx/html/*

# Stage 3: Copy your frontend files
COPY . /usr/share/nginx/html

# Port nginx serves on
EXPOSE 80
