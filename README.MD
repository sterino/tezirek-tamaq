# 🍽️ Tezirek Tamaq

**Tezirek Tamaq** is a microservice-based food delivery platform that connects users with restaurants.  
It’s built with **Golang**, **gRPC**, **PostgreSQL**, **Gin**, **Docker**, and **Swagger**.

---

## 🧩 Architecture Overview

The system is split into the following services:

### 🔐 Auth Service
- Handles user authentication.
- Generates and validates JWT tokens.

### 👤 User Service
- Manages user data (registration, retrieval).
- Supports role-based users (e.g. customer).

### 🍽️ Restaurant Service
- Manages restaurant records (CRUD).
- Associates restaurants with their orders.

### 🛒 Order Service
- Handles order creation, updates, deletion.
- Connects users and restaurants through orders.

### 🌐 API Gateway
- HTTP gateway using **Gin** framework.
- Validates JWT tokens.
- Routes requests to gRPC services.
- Includes Swagger UI at: `http://localhost:8080/swagger/index.html`

---

## ⚙️ Getting Started

### 📁 Clone the Repository

```bash
git clone https://github.com/shynkaaaa/tezirek_tamaq.git
cd tezirek_tamaq
```
## ⚙️ Environment Configuration
### Create a .env file and add the necessary variables

## 🚀  Run with Docker Compose
```bash
docker-compose up --build
```
### - This will automatically start:
- PostgreSQL
- All gRPC services (user, restaurant, order, auth)
- API Gateway with an HTTP interface

## 🌐 Доступные эндпоинты
| Сервис              | Порт  | Swagger UI            |
| ------------------- |-------| --------------------- |
| API Gateway         | 8080  | `/swagger/index.html` |
| User Service (gRPC) | 50051 |                       |
| Restaurant Service  | 50053 |                       |
| Order Service       | 50052 |                       |
| Auth Service        | 50051 |                       |

## 🔐 Authentication
### - Uses JWT (JSON Web Token)
### - The token must be included in the request header:
```bash
Authorization: Bearer <your_token>
```

## 📚 API Examples (via Swagger)
### - POST /auth/register – user registration
### - POST /auth/login – login and receive token

### - GET /user/{id} – fetch user data

### - POST /restaurant/create – create a new restaurant

### -  POST /order/create – place an order

## 🛠️ Technologies Used
- Golang
- gRPC

- Protocol Buffers

- Gin Web Framework

- Swagger (swaggo)

- PostgreSQL

- Docker & Docker Compose




