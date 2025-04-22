# 📌 URL Shortener Microservices

This is a practice project to learn Go by building a simple URL shortener using a microservices architecture.

## 🧠 Design Approach

This project is built with a focus on clean architecture principles:

- **Domain-Driven Design (DDD)** — The system is structured around the domain model to clearly separate core business logic from infrastructure concerns.

- **CQRS (Command Query Responsibility Segregation)** — Commands (writes) and queries (reads) are handled separately to improve scalability and maintainability.


## 👁️‍🗨️ Services Overview
This system is composed of three main microservices:

- **URL Shortener Service** — Responsible for generating and storing shortened URLs. It handles command-side operations (writes).

- **Redirector Service** — Resolves shortened URLs and redirects users to the original URL. It handles query-side operations (reads).
- **Statistics Service** — Collects usage data (e.g. redirects, user agents, geo info) and pushes it to a data warehouse for analytics.