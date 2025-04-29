#!/bin/bash

NAME="word2wl"

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Installing Go..."
    # Install Go (Ubuntu/Debian example)
    sudo apt update
    sudo apt install -y golang-go
fi

# Clone the repository
echo "Cloning the repository..."
git clone https://github.com/unsubble/${NAME}.git
cd ${NAME}

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Build the project
echo "Building the project..."
go build -o ${NAME}

# Make binary executable
chmod +x ${NAME}

# Move binary to a directory in PATH
sudo mv ${NAME} /usr/local/bin/

echo "Do you want to delete the leftover files? (y/n)"
read -r delete_files

if [[ "$delete_files" == "y" || "$delete_files" == "Y" ]]; then
    cd ..
    rm -rf ${NAME}
    echo "Leftover files deleted."
else
    echo "Leftover files retained."
fi

echo -n "Installation complete! You can now run the project with '${NAME}'..."
read
