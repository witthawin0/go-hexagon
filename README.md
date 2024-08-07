# E-Commerce Backend

This is an e-commerce backend project written in Golang, following the principles of hexagonal architecture (also known as ports and adapters architecture). This architecture helps in creating a maintainable, testable, and scalable application.

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Architecture](#architecture)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This e-commerce backend provides a robust platform for managing products, orders, customers, and payments. It is designed with a focus on modularity and separation of concerns to ensure easy maintenance and scalability.

## Features

- User Authentication and Authorization
- Product Management
- Order Management
- Customer Management
- Payment Integration
- RESTful API

## Architecture

The project follows the hexagonal architecture pattern, which separates the core business logic from the infrastructure and external services. This allows for easier testing and maintenance.

### Layers

1. **Domain Layer**: Contains the business logic and domain entities.
2. **Application Layer**: Contains the use cases and application services.
3. **Ports**: Interfaces that define the contracts for the application.
4. **Adapters**: Implementations of the ports, interacting with external systems such as databases and APIs.

### Directory Structure

```plaintext
.
├── cmd
│   └── main.go           # Entry point of the application
├── internal
│   ├── adapters          # Adapters for external systems (DB, API)
│   ├── application       # Application services and use cases
│   ├── domain            # Domain entities and business logic
│   ├── ports             # Port interfaces
│   └── infrastructure    # Infrastructure-related code (config, middleware)
├── pkg
│   └── utils             # Utility packages
└── README.md
