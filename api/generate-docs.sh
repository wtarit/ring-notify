#!/bin/bash

# Generate Swagger documentation
echo "Generating Swagger documentation..."

# Check if swag is installed
if ! command -v ~/go/bin/swag &> /dev/null; then
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate docs
~/go/bin/swag init

echo "Swagger documentation generated successfully!"
echo "You can view it at: http://localhost:1323/swagger/index.html"
