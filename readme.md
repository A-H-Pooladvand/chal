# Challenge

Description...

## Features

- **Database**: Utilizes PostgreSQL with `GORM` as the ORM for database operations.
- **Logging**: Offers logging capabilities to `file`, and `standard output (stdout) using Zap`.
- **Framework**: Built on top of `Gin`, a high-performance Go web framework.
- **Command Line Interface**: Employs `Cobra` for building command-line interfaces.
- **Commands**:
    - `serve`: Command to serve the microservice.
    - `migrate`: Command to perform database migrations.

### Why using godotenv package over Viper?
Secrets should be managed with a secret service like `hashicorp/vault`,
so I just used the simplest solution

### Why there is no unit testing?
Had no time in the time frame so...

## Usage

To use the app, follow these steps:

1. Hopefully with running `docker-compose up` you're good to go