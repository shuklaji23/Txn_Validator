# Transaction Validator

The Transaction Validator is a Go project that provides a framework for validating and processing transactions. It allows you to define rules and logic for validating transaction data and performs the necessary checks to ensure data integrity and compliance.

## Features

- JSON transaction data input
- Flexible rule-based validation
- Error reporting and handling
- Customizable transaction processing

## Installation

1. Ensure you have Go installed (version 1.16 or higher).
2. Clone the repository or download the source code.
3. Navigate to the project directory.

## Usage

1. Define your transaction validation rules in the `txnChecker.go` file.
2. Implement the transaction processing logic in the `routes_codes.go` file.
3. Customize the database reporting and handling in the `database_codes.go` file.
4. Build and run the project using the following command:

   ```shell
   go run main.go
   
## Configurations

The project provides a few configuration options that can be modified in the config.go file:

1. BlockCapacity: The maximum allowed size of blocks displayed in one go.
2. LocalHostPath: The local path can be set by the user themselves.
3. Now (variable; set to 5secs): The duration after which a transaction validation times out.
   
Feel free to adjust these values according to your project requirements.
