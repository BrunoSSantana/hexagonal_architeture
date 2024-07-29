
# Hexagonal Architecture

This project is designed to explore hexagonal architecture using Golang, drawing on the content presented in FullCycle's course on hexagonal architecture, following the exposed content.

## How to Use

To set up and run the web server:

**1. Clone the Project**
```bash
git clone https://github.com/BrunoSSantana/hexagonal_architeture.git
```

Ensure Docker and Docker Compose are installed on your machine.

**2. Start the Containers**
```bash
docker compose up -d && docker exec -it app-product bash
```

**3. Run the Server Inside the Container**
```bash
go run main.go http
```

For command-line interaction and to view available options, run:
```bash
go run main.go cli -h
```
