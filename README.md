# Fraud Auth Service

## 🚀 Overview

The Fraud Auth Service is responsible for authentication and authorization in the Real-Time Fraud Detection Platform.

It provides secure user registration and login functionality using JWT-based authentication.

This service acts as the identity provider for the entire distributed system.

---

## 🎯 Purpose in Overall Architecture

In a distributed system, authentication must be isolated from business logic services.

This service:

- Issues JWT tokens
- Validates user credentials
- Provides secure identity to downstream services
- Enables API Gateway to enforce authorization

Without this service:
- No secure access control
- No user identity propagation
- No rate limiting per user
- No traceability in fraud detection

---

## 🧱 Responsibilities

- User signup
- User login
- Password hashing (bcrypt)
- JWT generation
- Rate limiting (Redis - future)
- Token validation middleware (future)

---

## 🗄 Database

### PostgreSQL

Table:
- users

Reason:
Authentication requires strong consistency and ACID guarantees.

---

## 🔐 Security Strategy

- Password hashing using bcrypt
- JWT tokens with expiration
- Centralized token validation
- No plaintext password storage

---

## 🔄 Interaction with Other Services

1. Client logs in via API Gateway.
2. Gateway forwards request to Auth Service.
3. Auth Service returns JWT.
4. Gateway validates JWT on every protected route.
5. Downstream services trust user identity from JWT claims.

---

## 📦 Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- JWT
- bcrypt
- Redis (rate limiting - upcoming)

---

## 🧠 Distributed System Role

In distributed architectures:

- Authentication must be stateless.
- Services should not manage sessions.
- JWT enables horizontal scaling.

This service is designed to be fully stateless and horizontally scalable.

---

## 📈 Future Enhancements

- Refresh tokens
- OAuth integration
- Role-based access control (RBAC)
- Redis-backed token blacklist
- Rate limiting middleware