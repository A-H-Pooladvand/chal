# Challenge

Description...

## Features

- **Life cycle**: Utilizes Uber FX. Uber FX is a dependency injection framework for Go that simplifies dependency management, promotes code reusability, and allows for implementing graceful shutdowns.
- **Database**: Utilizes PostgreSQL with `GORM` as the ORM for database operations.
- **Logging**: Offers logging capabilities to `file`, and `standard output (stdout) using Zap`.
- **Framework**: Built on top of `Gin`, a high-performance Go web framework.
- **Command Line Interface**: Employs `Cobra` for building command-line interfaces.
- **Commands**:
    - `serve`: Command to serve the microservice.
    - `migrate`: Command to perform database migrations.

### Why uber FX
- Manages Dependencies: Fx simplifies dependency injection.
- Reduces Globals: Fx encourages you to remove global variables and functions from your code.
- Promotes Reusability: By making dependencies explicit, Fx allows you to create loosely coupled components that can be easily reused in different parts of your application
- Testing Friendly: Fx allows you to easily define mock dependencies for unit testing.
- Lifecycle Hooks: Fx provides hooks like onStart and onStop that you can use to define logic for application startup and shutdown respectively.

### Why using godotenv package over Viper?
Secrets should be managed with a secret service like `hashicorp/vault`,
so I just used the simplest solution

### Why there is no unit testing?
Had no time in the time frame so...

## Usage

To use the app, follow these steps:

1. Hopefully with running `docker-compose up` you're good to go