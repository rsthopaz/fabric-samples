#!/bin/bash

# Add all changes
git add .

# Ask user for commit message
read -p "Enter commit message: " message

# Commit with the message
git commit -m "$message"

# Push to current branch
git push
