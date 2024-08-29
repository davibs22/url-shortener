# Url-Shortener

Url-Shortener is a simple and efficient URL shortener developed in Go using the Fiber library. This project allows users to shorten long URLs and redirect them to the original links quickly. The project also uses Redis for persistent storage of shortened URLs.

## Features

- **URL Shortening**: Convert long URLs into shorter versions.
- **Redirection**: Redirect shortened URLs to the original address.
- **Redis Storage**: Shortened URLs are stored in Redis for persistence.

## Technologies Used

- **Go**: v1.23.0
- **Fiber**: Web framework for Go, inspired by Express.js.
- **Redis**: Used for storing and retrieving shortened URLs.

## How to Use

### Prerequisites

- [Go](https://golang.org/dl/) v1.23.0 or later installed.
- [Redis](https://redis.io/download) installed and running.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/davibs22/url-shortener
    cd url-shortener
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the Redis server (if not already running):

    ```bash
    redis-server
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

### Endpoints

- `POST /shorten`: Receives a long URL and returns a shortened version.
- `GET /:hash`: Redirects the shortened URL to the original URL.
- `DELETE /shorten`: Delete the shortened URL.

### Usage Example

1. **Shorten a URL**:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"url": "http://localhost"}' http://localhost:3000/shorten
    ```

2. **Redirection**:

    Visit `http://localhost:3000/16220` in your browser to be redirected to the original URL.

## Contribution

Contributions are welcome! Feel free to open issues or pull requests to improve this project.

## License

This project is licensed under the [MIT License](LICENSE).
