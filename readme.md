# Go Figure API

A Go backend service for the go figure [web app](https://github.com/ImparaAI/go-figure-web). This is inspired by the [Fourier series circle drawings video by 3Blue1Brown](https://www.youtube.com/watch?v=r6sGWTCMz2k).

This application accepts a series of continuous points captured at even time intervals, and a requested number of draw vectors to calculate. The calculation job is put onto a queue that is processed continuously. Once the job is complete, the results are available at the submission's uri as a list of vectors with *real* and *imaginary* parts that can then be animated.

[![Build Status](https://travis-ci.org/ImparaAI/go-figure-api.png?branch=master)](https://travis-ci.org/ImparaAI/go-figure-api)

## Routes

### /submission
Method: `POST`

Input:

```json
{
	"drawVectorCount": 20,
	"points": [
		{"x": 143, "y": 45, "time": 0},
		{"x": 144, "y": 49, "time": 0.01}
	]
}
```

Output:

```json
{"id": 1}
```

### /submission/:id
Method: `GET`

Output:

```json
{
	"id": 1,
	"points": [
		{"x": 143, "y": 45, "time": 0},
		{"x": 144, "y": 49, "time": 0.01}
	],
	"drawVectors": {
		"calculated": [
			{"n": 0, "real": 0, "imaginary": 2},
			{"n": 1, "real": 0, "imaginary": 2.5},
			{"n": -1, "real": 0.5, "imaginary": 2.7}
		],
	},
	"dateCreated": "2020-01-01 12:00:00"
}
```

### /submission/:id/draw-vectors/:int
Method: `POST`

Output: `int` *(Current requested count)*

## Tests

Run `go test ./...`.