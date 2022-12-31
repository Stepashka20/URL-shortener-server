# URL Shortener Server

A simple URL shortening service that allows you to create short, memorable links for long URLs.


## Frontend
The frontend for this project can be found [here](https://github.com/Stepashka20/URL-shortener)

## Features
- [x] Generate short links from long URLs
- [ ] Customize the short link with your own desired text
- [ ] View click statistics for each short link

## Requirements
- Go
- MongoDB

## Installation
1. Clone the repository: `git clone https://github.com/Stepashka20/URL-shortener-server.git`
2. Install dependencies: `go get`
3. In MongoDB create a new db `urls` and collection `shorturls` 
4. Copy `.env.example` to `.env` and update the database connection details
5. Start the server: `go run main.go`

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.