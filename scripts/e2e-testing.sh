#!/bin/bash

# Base URL
API_URL=http://localhost:8080


# Introduction to the script.
echo "Welcome to 'seed-rest-api' application!"
echo "Before running the end-to-end tests, please ensure that you have run 'make start'!"; echo

# Testing '/api/v1'.
echo
echo "Running end-to-end testing..."
echo "Testing GET route '/api/v1'..."
curl $API_URL/api/v1; echo

# Testing '/api/v1/users'.
echo
echo "Testing GET route '/api/v1/users'..."
curl $API_URL/api/v1/users; echo
echo
echo "Testing POST route '/api/v1/users'..."
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Lucy Heartfilia","address":"Shinhotaka, Japan"}' $API_URL/api/v1/users; echo
