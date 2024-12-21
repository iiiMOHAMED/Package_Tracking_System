# Step 1: Use a Node.js image for building the Angular app
FROM node:18-alpine AS build

# Set the working directory
WORKDIR /app

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Angular app for production
RUN npm run build --prod

# Step 2: Use a minimal NGINX image to serve the built files
FROM nginx:alpine

# Switch to root to handle permission-sensitive operations
USER root

# Create necessary directories with permissions
RUN mkdir -p /var/cache/nginx /var/log/nginx /usr/share/nginx/html && \
    chmod -R 755 /var/cache/nginx /var/log/nginx /usr/share/nginx/html && \
    chown -R nginx:nginx /var/cache/nginx /var/log/nginx /usr/share/nginx/html

# Copy the NGINX configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the built Angular app from the previous stage to the NGINX web directory
COPY --from=build /app/dist/shipping-frontend /usr/share/nginx/html

# Expose the port for the web server
EXPOSE 80

# Start NGINX to serve the app
CMD ["nginx", "-g", "daemon off;"]
