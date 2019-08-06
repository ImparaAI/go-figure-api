# Go Figure API

A Go backend service for the go figure [web app](https://github.com/ImparaAI/go-figure-web). This is inspired by the [Fourier series circle drawings video by 3Blue1Brown](https://www.youtube.com/watch?v=r6sGWTCMz2k).

This application accepts a series of continuous points captured at even time intervals, and a requested number of draw vectors to calculate. The calculation job is put onto a queue that is processed continuously. Once the job is complete, the results are available at the drawing's uri as a list of vectors with *real* and *imaginary* parts that can then be animated.

[![Build Status](https://travis-ci.org/ImparaAI/go-figure-api.png?branch=master)](https://travis-ci.org/ImparaAI/go-figure-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/ImparaAI/go-figure-api)](https://goreportcard.com/report/github.com/ImparaAI/go-figure-api)


## Routes

### /drawing
Method: `POST`

Input:

```json
{
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

### /drawing/:id
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

### /drawings/recent
Method: `GET`

Output:

```json
[
	{
		"id": 2,
		"svgPath": "M 364 113 L 364 113 L 357 113 L 348 113 L 342 113 L 337 113 L 333..."
	},
	{
		"id": 1,
		"svgPath": "M 364 113 L 364 113 L 357 113 L 348 113 L 342 113 L 337 113 L 333..."
	}
]
```

## Tests

Run `go test ./.../test`.