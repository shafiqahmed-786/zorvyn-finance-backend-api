# 🚀 Finance Backend API — Zorvyn Backend Developer Internship Assignment

A production-grade **FinTech backend service** built in **Go + Fiber + PostgreSQL + GORM**, designed for secure financial record management, RBAC-based access control, audit logging, analytics dashboards, and API-first developer experience.

This project was built as part of the **Backend Developer Intern Screening Assignment for Zorvyn** and aims to reflect **real-world backend engineering standards**.

---

# ✨ Features

## 🔐 Authentication & Authorization
- JWT-based authentication
- Secure password hashing
- Role-Based Access Control (**Admin / Analyst / Viewer**)
- Protected route middleware
- Bearer token support in Swagger UI

---

## 👥 User Management
- User registration
- User login
- User listing
- Role assignment
- Active user status
- UUID-based user IDs

---

## 💰 Financial Records
- Create financial records
- List all records
- Soft delete support
- Record filtering
- Pagination support
- Income / Expense classification
- Category-wise storage
- Notes and metadata

---

## 📊 Dashboard Analytics
- Total income
- Total expenses
- Net balance
- Recent transaction count
- Category breakdown
- Monthly trend-ready analytics structure
- Recent activity support

---

## 📝 Audit Logs
- Financial record activity tracking
- Secure audit trail support
- Designed for compliance-friendly fintech workflows

---

## 📚 API Documentation
- Interactive Swagger UI
- Request/response schemas
- DTO + model documentation
- JWT authorize support
- Ready for frontend integration

---

# 🏗️ Tech Stack

- **Go**
- **Fiber**
- **PostgreSQL**
- **GORM**
- **JWT**
- **Docker**
- **Swagger**
- **UUID / pgcrypto**

---

# 📂 Project Structure

```text
finance-backend/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── seed/
│       └── main.go
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── dto/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   ├── services/
│   └── utils/
│
├── docker-compose.yml
├── Makefile
├── .env
└── README.md

⚙️ Setup Instructions
1️⃣ Clone
git clone <your-repo-url>
cd finance-backend
2️⃣ Start PostgreSQL
docker-compose up -d
3️⃣ Configure environment

Create .env

PORT=8082
DB_HOST=127.0.0.1
DB_PORT=5435
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=finance_db
JWT_SECRET=supersecretkey
4️⃣ Enable UUID extension
docker exec -it finance_postgres psql -U postgres
\c finance_db
CREATE EXTENSION IF NOT EXISTS pgcrypto;
\q
5️⃣ Run server
go run cmd/server/main.go
6️⃣ Generate Swagger docs
swag init -g server/main.go -d cmd,internal --parseDependency --parseInternal
7️⃣ Seed demo data
go run cmd/seed/main.go
🌐 API Documentation

Swagger UI available at:

http://localhost:8082/swagger/index.html