#!/bin/bash
echo "Installing Focus-Symphony..."
go build -o focus-symphony main.go
mkdir -p ~/.local/bin
mv focus-symphony ~/.local/bin/
echo "Done! Run 'focus-symphony' from anywhere."
