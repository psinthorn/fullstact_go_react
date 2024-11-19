# fullstact_go_react

Fullstack web application using Golang as API backend and html and react as frontend

This project is a fullstack web application using Golang as API backend and html, react, nextjs and angular as frontend. The project is a simple CRUD application that allows users to create, read, update and delete items from a list. The project is a good starting point for anyone looking to learn how to build a fullstack web application using Golang and React.

## Getting Started

### Prerequisites

#### Golang

You need to have Golang installed on your machine. You can download and install Golang from the official website: https://golang.org/dl/

#### Node.js

You need to have Node.js installed on your machine. You can download and install Node.js from the official website: https://nodejs.org/en/download/

#### Yarn or NPM

You need to have Yarn or NPM installed on your machine. You can download and install Yarn from the official website: https://yarnpkg.com/getting-started/install or NPM from the official website: https://www.npmjs.com/get-npm

### Installation

1. Clone the repository

```sh
git clone repo_url
```

2. Change directory to the project directory

```sh
cd fullstack_go_react
```

3. Install the dependencies

```sh
cd frontend
npm install
```

```sh
cd backend
go mod download
```

4. Start the frontend

```sh
cd frontend
npm run dev
```

5. Start the backend

```sh
cd backend
go run main.go
```

6. Open your browser and navigate to http://localhost:3000

## Usage

### This project is use Dokku for deployment to Digital Ocean server and use Postgres as database server and use Nginx as reverse proxy server and use Let's Encrypt for SSL certificate and use Cloudflare for DNS server.

### Dockerfile and docker-compose.yml are included for local development and testing.

### Makefile is included for easy deployment testing and development.

## Built With

- [Golang](https://golang.org/) - The programming language used
- [React](https://reactjs.org/) - The frontend library used
- [Next.js](https://nextjs.org/) - The frontend framework used
- [Angular](https://angular.io/) - The frontend framework used
- [Node.js](https://nodejs.org/) - The runtime environment used
- [Yarn](https://yarnpkg.com/) - The package manager used
- [NPM](https://www.npmjs.com/) - The package manager used
- Tailwind CSS - The CSS framework used
