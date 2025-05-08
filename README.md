# Pinned Menu 🚀

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-menu)](https://goreportcard.com/report/github.com/patricksferraz/pinned-menu)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-menu?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-menu)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, high-performance menu management system built with Go, featuring a robust API and scalable architecture.

## ✨ Features

- 🚀 High-performance API built with Fiber
- 📊 PostgreSQL and SQLite database support
- 🔄 Real-time updates with Kafka integration
- 📝 Swagger API documentation
- 🐳 Docker and Kubernetes support
- 🔐 Environment-based configuration
- 🧪 Comprehensive testing suite

## 🛠️ Tech Stack

- **Backend**: Go 1.18+
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL, SQLite
- **ORM**: GORM
- **Message Broker**: Kafka
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Documentation**: Swagger
- **Testing**: Go testing framework

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose
- PostgreSQL (optional, SQLite is available for development)
- Kafka (optional, for real-time features)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/pinned-menu.git
cd pinned-menu
```

2. Copy the environment file and configure it:
```bash
cp .env.example .env
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
make run
```

### Docker Deployment

```bash
docker-compose up -d
```

## 📚 API Documentation

Once the application is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/
```

## 🏗️ Project Structure

```
.
├── app/          # Application layer
├── cmd/          # Command line interface
├── domain/       # Domain models and interfaces
├── infra/        # Infrastructure implementations
├── k8s/          # Kubernetes configurations
└── utils/        # Utility functions
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Patrick Ferraz - Initial work - [GitHub](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/go-gorm/gorm)
- [Swagger](https://github.com/swaggo/swag)
- And all other amazing open-source projects that made this possible!

---

⭐️ If you like this project, please give it a star on GitHub!
