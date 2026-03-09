# Transactions API

## Overview

The **Transactions API** is a demonstration RESTful web service built in Go, designed to manage user accounts and process various financial transactions. This project showcases best practices in software architecture, including clean architecture principles, dependency injection, and the strategy design pattern for handling different transaction types.

The API supports core banking-like operations such as account creation, retrieval, and transaction processing (e.g., purchases, withdrawals). It integrates with a MySQL database for persistent storage and uses modular components to separate concerns like data access, business logic, and HTTP handling.

## Features

- **Account Management**:
  - Create new accounts with a document number.
  - Retrieve account details by ID.

- **Transaction Processing**:
  - Process transactions using the Strategy Pattern, allowing easy extension for new transaction types.
  - Supported transaction types:
    - **Normal Purchase** (Operation Type ID: 1): Standard debit transaction.
    - **Purchase with Installments** (Operation Type ID: 2): Debit with installment logic (simplified).
    - **Withdrawal** (Operation Type ID: 3): Debit for cash withdrawal.
    - **Credit Voucher** (Operation Type ID: 4): Credit transaction to add funds.

- **Validation**:
  - Input validation for account and transaction requests to ensure data integrity.

- **Database Integration**:
  - MySQL database with repository pattern for data access.
  - Models for Account and Transaction entities.

- **RESTful API**:
  - Standard HTTP methods (GET, POST) for endpoints.
  - JSON request/response format.

- **Middleware**:
  - Logging middleware to track incoming requests.

- **Modular Architecture**:
  - Separation of concerns with packages for config, handlers, services, etc.
  - Dependency injection via a custom container for testability and flexibility.

- **TestCases**:
  - TestCase are made available for all the possibilities where business logic might go wrong.
  - For both the flows accounts and transactions.

## Tech Stack

- **Programming Language**: Go (version 1.19 or later)
- **Database**: MySQL (for data persistence)
- **Web Framework**: Standard library (`net/http`) for HTTP server and routing
- **Architecture Patterns**:
  - Clean Architecture: Layers for entities, use cases, interfaces, and frameworks.
  - Repository Pattern: Abstracts data access.
  - Strategy Pattern: Handles different transaction types polymorphically.
- **Dependency Injection**: Custom container for managing service dependencies.
- **Other Tools**: Environment variables for configuration, standard logging.

## Project Structure

The project follows a clean, modular structure to promote maintainability:

```
transaction/
├── config/           # Loads configuration from environment variables (e.g., DB credentials, port).
├── container/        # Dependency injection container to wire up services and repositories.
├── db/               # Database connection setup, initialization and seed data file.
├── dto/              # Data Transfer Objects (DTOs) for API requests and responses (e.g., AccountRequest, TransactionResponse).
├── handlers/         # HTTP request handlers that process API endpoints and delegate to services.
├── middleware/       # HTTP middleware, such as logging for request tracking.
├── models/           # Database models representing entities (e.g., Account, Transaction).
├── repository/       # Data access layer with interfaces and implementations for database operations.
├── routes/           # Route definitions and registration with the HTTP server.
├── services/         # Business logic layer containing use cases (e.g., account creation, transaction processing).
├── strategies/       # Strategy implementations for different transaction types.
├── validator/        # Validation logic for incoming requests.
├── tests/            # Unit TestCases for account and transaction business core logics.
├── main.go           # Application entry point that sets up the server and starts listening.
└── README.md         # This documentation file.
```

## Installation and Setup

Follow these steps to set up and run the Transactions API on your local machine.

### Prerequisites

- **Go**: Install Go 1.19 or later from [golang.org](https://golang.org/dl/).
- **MySQL**: Install MySQL Server (version 5.7 or later). You can download it from [mysql.com](https://dev.mysql.com/downloads/mysql/).
- **Git**: For cloning the repository.
- **IDE/Editor**: Visual Studio Code or any Go-compatible editor (optional but recommended).

### Step-by-Step Installation

1. **Clone the Repository**:
   Open your terminal and run:
   ```bash
   git clone <repository-url>  # Replace with the actual Git repository URL
   cd transaction
   ```

2. **Install Dependencies**:
   Ensure you have Go modules enabled. Run:
   ```bash
   go mod tidy
   ```
   This will download and install all required Go packages.

3. **Set Up the Database**:
   - Create a new MySQL database (e.g., named `transactions_db`).
   - Update the database credentials in `config/config.go` or set environment variables (see below).
   - Manually create the necessary tables based on the models in `models/`. For example:
     ```sql
     CREATE TABLE accounts (
         id INT AUTO_INCREMENT PRIMARY KEY,
         document_number VARCHAR(255) NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     );
     CREATE TABLE transactions (
         id INT AUTO_INCREMENT PRIMARY KEY,
         account_id INT NOT NULL,
         amount DECIMAL(10,2) NOT NULL,
         transaction_type_id INT NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
         FOREIGN KEY (account_id) REFERENCES accounts(id)
     );
     CREATE TABLE transactions_types (
         id INT AUTO_INCREMENT PRIMARY KEY,
         description VARCHAR(255) NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     );
     ```
   - After tables have been created properly, also please do run seed.sql file to pre-seed the data into database tables.

     **Note**: In a real project, use migrations (e.g., via `golang-migrate`).

4. **Configure Environment Variables**:
   Create a `.env` file in the root directory or export variables in your shell. Defaults are provided in the code:
   - `PORT=8080`  # Server port
   - `ENABLE_REST_API=true`  # Enable API endpoints
   - `DB_HOST=localhost`  # MySQL host
   - `DB_PORT=3306`  # MySQL port
   - `DB_USER=root`  # MySQL username
   - `DB_PASSWORD=password`  # MySQL password
   - `DB_NAME=transactions_db`  # Database name

5. **Build and Run the Application**:
   - Build the project:
     ```bash
     go build -o transaction main.go
     ```
   - Run the application:
     ```bash
     ./transaction
     ```
   The server will start on `http://localhost:8080` (or the configured port).

### Troubleshooting

- If you encounter database connection errors, verify your MySQL server is running and credentials are correct.
- Ensure the database tables are created as described.
- Check Go version with `go version` to confirm compatibility.

## Usage

Once the server is running, interact with the API using tools like Postman, curl, or any HTTP client. All requests and responses use JSON format.

### API Endpoints

#### 1. Create Account
- **Endpoint**: `POST /accounts`
- **Description**: Creates a new account.
- **Request Body** (JSON):
  ```json
  {
    "document_number": "123456789"
  }
  ```
- **Response** (JSON):
  ```json
  {
    "id": 1,
    "document_number": "123456789"
  }
  ```
- **Status Codes**:
  - 201: Created
  - 400: Bad Request (invalid input)
  - 500: Internal Server Error

#### 2. Get Account
- **Endpoint**: `GET /accounts/{accountId}`
- **Description**: Retrieves details of an account by ID.
- **Response** (JSON):
  ```json
  {
    "id": 1,
    "document_number": "123456789"
  }
  ```
- **Status Codes**:
  - 200: OK
  - 404: Not Found
  - 500: Internal Server Error

#### 3. Create Transaction
- **Endpoint**: `POST /transactions`
- **Description**: Processes a transaction for an account.
- **Request Body** (JSON):
  ```json
  {
    "account_id": "1",
    "amount": 100.50,
    "operation_type_id": "1"
  }
  ```
- **Response** (JSON):
  ```json
  {
    "id": 1,
    "account_id": 1,
    "amount": 100.50,
    "operation_type_id": 1,
    "created_at": "2023-10-01T12:00:00Z"
  }
  ```
- **Status Codes**:
  - 201: Created
  - 400: Bad Request (invalid input or insufficient funds)
  - 404: Account Not Found
  - 500: Internal Server Error

### Example Usage with curl

- **Create an Account**:
  ```bash
  curl -X POST http://localhost:8080/accounts \
       -H "Content-Type: application/json" \
       -d '{"document_number": "123456789"}'
  ```

- **Get an Account**:
  ```bash
  curl -X GET http://localhost:8080/accounts/1
  ```

- **Create a Transaction**:
  ```bash
  curl -X POST http://localhost:8080/transactions \
       -H "Content-Type: application/json" \
       -d '{"account_id": "1", "amount": 100.50, "operation_type_id": "1"}'
  ```

### Testing the API

- Use Postman to send requests and inspect responses.
- Check server logs in the terminal for request details (via middleware).
- For automated testing, add unit tests (see Development section).

## Development

### Running Tests
- Add unit tests in files like `accounts_test.go`, `transaction_test.go`, etc.
- Run tests:
  ```bash
  go test ./...
  ```

### Building for Production
- Build a binary:
  ```bash
  go build -o transaction main.go
  ```
- For cross-compilation (e.g., for Linux):
  ```bash
  GOOS=linux GOARCH=amd64 go build -o transaction main.go
  ```

### Extending the Project
- **Add New Transaction Types**: Implement a new strategy in `strategies/` and register it in the container.
- **Add Authentication**: Integrate JWT or OAuth in middleware.
- **Improve Error Handling**: Use custom error types and HTTP status codes.

## Contributing

We welcome contributions to improve this demo project!

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/new-feature`).
3. Make changes and add tests.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/new-feature`).
6. Create a Pull Request.

Please follow Go conventions and ensure tests pass.

## License

This project is provided for demonstration purposes only. It is not licensed for commercial use. For any production deployment, consider applying an appropriate open-source license (e.g., MIT or Apache 2.0).

## Contact

For questions or feedback, please open an issue in the repository.