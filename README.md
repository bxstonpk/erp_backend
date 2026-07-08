# ERP Backend

A scalable Enterprise Resource Planning (ERP) backend developed in **Go (Golang)** following Clean Architecture and Domain-Driven Design (DDD) principles.

This project is designed to provide a modular and maintainable ERP platform that can be extended continuously as business requirements evolve.

---

## 🚀 Current Modules

### Authentication
- JWT Authentication
- Login / Logout
- Refresh Token
- Password Encryption
- Role-Based Authorization (RBAC)
- Permission Management

### User Management
- User CRUD
- Role Assignment
- User Profile
- Account Status

### Sales
- Customer Management
- Sales Quotation
- Sales Order
- Invoice
- Sales History

---

## 🛠 Tech Stack

- Go (Golang)
- Gin / Gorilla Mux *(depending on implementation)*
- GORM
- MySQL
- JWT
- Docker
- Docker Compose

---

## 🏗 Architecture

This project follows:

- Clean Architecture
- Domain-Driven Design (DDD)
- Repository Pattern
- Dependency Injection
- RESTful API

Project structure:

```
.
├── cmd/
├── internal/
│   ├── auth/
│   ├── user/
│   ├── sale/
│   ├── shared/
│   └── infrastructure/
├── pkg/
├── configs/
├── docs/
└── migrations/
```

---

## 📦 Features

- Modular Design
- REST API
- JWT Authentication
- RBAC Authorization
- Database Migration
- Validation
- Logging
- Configuration Management
- Error Handling

---

## 📅 Roadmap

Current development focuses on building the core ERP foundation.

### Phase 1
- [x] Authentication
- [x] User Management
- [ ] Sales Module

### Phase 2
- [ ] Customer Relationship Management (CRM)
- [ ] Purchasing
- [ ] Inventory Management
- [ ] Warehouse Management
- [ ] Product Management

### Phase 3
- [ ] Accounting
- [ ] Finance
- [ ] General Ledger
- [ ] Accounts Payable
- [ ] Accounts Receivable
- [ ] Tax Management

### Phase 4
- [ ] Production Planning
- [ ] Material Requirement Planning (MRP)
- [ ] Manufacturing Execution
- [ ] Quality Control

### Phase 5
- [ ] Human Resources (HR)
- [ ] Payroll
- [ ] Leave Management
- [ ] Performance Evaluation

### Phase 6
- [ ] Dashboard & Analytics
- [ ] Business Intelligence
- [ ] Notification Service
- [ ] Report Generator
- [ ] Audit Log
- [ ] API Gateway
- [ ] Multi-Tenant Support

---

## ⚙️ Getting Started

### Clone repository

```bash
git clone https://github.com/yourname/erp-backend.git
cd erp-backend
```

### Install dependencies

```bash
go mod tidy
```

### Configure environment

Create `.env`

```env
APP_NAME=ERP Backend

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=erp

JWT_SECRET=your-secret
```

### Run

```bash
go run cmd/main.go
```

or

```bash
docker compose up --build
```

---

## 📖 API Documentation

Coming Soon

- Swagger/OpenAPI

---

## 🧪 Testing

```bash
go test ./...
```

---

## 🤝 Contributing

Contributions, suggestions, and issue reports are welcome.

Feel free to open an Issue or Pull Request.

---

## 📌 Project Status

🚧 **Work In Progress**

This ERP backend is under active development. New modules and improvements will be added continuously.

---

## 📄 License

MIT License