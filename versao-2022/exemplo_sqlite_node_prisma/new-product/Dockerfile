# ./new-product/Dockerfile

FROM node:18

# Create app directory
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy application source code
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Start the app
CMD [ "node", "server.js" ]
