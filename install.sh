#!/bin/bash

INSTALL_PATH="${1:-/usr/local/bin}"

echo "🔨 Building work..."
go build -o work

if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi

echo "📦 Installing to $INSTALL_PATH..."
cp ./work "$INSTALL_PATH/"

if [ $? -eq 0 ]; then
    echo "🎉 Installed work successfully to $INSTALL_PATH"
else
    echo "❌ Installation failed. You may need to run with sudo:"
    echo "sudo ./install.sh"
fi
