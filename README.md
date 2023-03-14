# URL Shortener Project

This project is a URL shortener application built with Go Fiber, Redis, and Docker. The application allows users to shorten long URLs into short, easy-to-remember links.

## Features

- Shorten long URLs into short, easy-to-remember links
- Redirect shortened links to their original URLs
- Can use the same shorten URL for maximum 24 hours

## Tech Stack

The tech stack used in this project includes:

- Go Fiber - A fast and lightweight web framework for Go.
- Redis - An open-source, in-memory data structure store.
- Docker - A containerization platform for developing, shipping, and running applications.

## Getting Started

### Prerequisites

To run this application, you will need to have the following installed on your machine:

- Docker
- Docker Compose

### Installation

To install and run the application:

1. Clone this repository to your local machine.

```
git clone https://github.com/ikhwankhaliddd/URL-Shorten-App.git
```


2. Change into the project directory.

```
cd url-shortener
```
 
3. Build and start the application using Docker Compose.
```
docker-compose up -d
```


4. Access the application at `http://localhost:3000`.

## Usage

The example request of the shorten URL API:

```
{
    "url" : "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    "short" : "TestShortenURL"
    "expiry" : "24"
}
```

You can customize the short URL and the expiration time. The default value of short URL is UUID and the default value of expiry is 24 hours.

To access the original URL:

1. Enter the shortened URL into your browser's address bar.
2. Hit enter.
3. You will be redirected to the original URL.

The example response of the shorten URL API:

```
{
    "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    "short": "localhost:3000/TEST",
    "expiry": 24,
    "rate_limit": 9,
    "rate_limit_reset": 30
}
```

## Acknowledgements

This application was inspired by `Bit.ly URL Shortener`.


