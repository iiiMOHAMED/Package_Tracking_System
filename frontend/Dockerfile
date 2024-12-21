# Step 1: Use a minimal Node.js image based on Alpine for building the Angular app
FROM node:18-alpine AS build

# Set the working directory
WORKDIR /app

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Install dependencies (from package.json and package-lock.json)
COPY package*.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Step 2: Expose the development server port and run Angular development server
EXPOSE 4200

# Start the Angular development server and bind to all network interfaces (0.0.0.0)
CMD ["npm", "run", "start", "--", "--host", "0.0.0.0", "--disable-host-check"]
