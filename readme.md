A Go backend service for the go figure [web app](https://github.com/ImparaAI/go-figure-web).

[![Build Status](https://travis-ci.org/ImparaAI/go-figure-backend.png?branch=master)](https://travis-ci.org/ImparaAI/go-figure-backend)

# Routes

## /submission
Method: `POST`

Input: `pixels=[int, int]`

## /submission/:id
Method: `GET`

Output: `{...}`