# Transfer Money API

This is the Transfer Money API, which allows users to perform various operations related to transferring money.

## Documentation

The documentation for the API endpoints can be found on Postman. Click the link below to access the documentation:

[Transfer Money API Documentation](https://documenter.getpostman.com/view/26374053/2s9Yyth1EF)

## Features

- User authentication (login, registration)
- Money transfer between accounts
- Deposit money into user accounts
- Retrieve user details and transaction history

## Getting Started

To get started with the Transfer Money API, follow these steps:
1. Clone this repository to your local machine.
2. Install the required dependencies using `go mod tidy`.
3. Configure the database connection in the `.env` file.
4. Run the API server using `go run main.go`.
5. Access the API endpoints using tools like Postman or curl.

## API Endpoints

- `POST /api/auth/login`: User login endpoint.
- `POST /api/auth/register`: User registration endpoint.
- `POST /api/auth/logout`: User logout endpoint.
- `POST /api/transaction/deposit`: Deposit money into user account.
- `POST /api/transaction/transfer`: Transfer money between accounts.
- `GET /api/users`: Get all users.
- `GET /api/users/:id`: Get user details by ID.
- `PUT /api/users/:id`: Update user details by ID.

## Technologies Used

- Go (Golang) - Programming language for the backend server.
- Gin - Web framework for building APIs in Go.
- GORM - ORM (Object-Relational Mapping) library for working with databases in Go.
- PostgreSQL - Database management system for storing user and transaction data.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
