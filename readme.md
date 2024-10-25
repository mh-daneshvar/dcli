# dcli: CLI Automation for Managing Multiple Project Services

## Overview

Managing multiple projects and their services simultaneously can be a hassle, especially when switching between them requires manually starting and stopping various prerequisites. **dcli** is a Go-based command-line tool designed to make it easier to switch between projects by providing an interactive, prompt-based interface that manages your Docker services and dependencies.

The tool works by reading a YAML configuration file where you define your projects, their services, dependencies, and Docker Compose files. Using the prompt UI, you can quickly select a project and manage its services directly from the terminal.

For a sample configuration and example services, please refer to the [dcli-other-projects repository](https://github.com/mh-daneshvar/dcli-other-projects), which contains several example project setups.

## Key Features

- **Service Management**: Define individual services and their respective Docker Compose files.
- **Dependency Resolution**: Automatically start dependencies in the correct order when starting a service.
- **Common Services**: Define shared services and containers that are used across multiple services in a project.
- **Docker Compose Integration**: Use Docker Compose files for automatic container management based on your project setup.
- **Interactive Prompt UI**: A user-friendly interface powered by `promptui` that allows users to select and manage services through terminal prompts.

## Project Repositories

- **Main Project (dcli)**: [https://github.com/mh-daneshvar/dcli](https://github.com/mh-daneshvar/dcli)
- **Sample Projects (dcli-other-projects)**: [https://github.com/mh-daneshvar/dcli-other-projects](https://github.com/mh-daneshvar/dcli-other-projects)

## Sample YAML Configuration

Below is an example YAML file that defines a project and its services. You can use this to configure your own project setup:

```yaml
projects:

  Project One:
    services:
      - label: Auth Service
        dependencies:
          - Users Service
        containers:
          - graylog
        docker_compose_file_path: /path/to/auth-service/docker-compose.yaml
      - label: Users Service
        docker_compose_file_path: /path/to/users-service/docker-compose.yaml
      - label: Catalogue Service
        docker_compose_file_path: /path/to/catalogue-service/docker-compose.yaml
    common:
      containers:
        - open-webui
      docker_compose_file_path: /path/to/common-services/docker-compose.yaml
```

### Key Components in YAML

- **projects**: This is the top-level key containing all your project definitions.
- **Project One**: This is a sample project containing multiple services.
- **services**: A list of services associated with the project.
   - **label**: The name of the service (e.g., "Auth Service", "Users Service").
   - **dependencies**: Other services that must be running for this service to start.
   - **containers**: Optional list of Docker containers that need to be running for the service.
   - **docker_compose_file_path**: Path to the Docker Compose file responsible for managing the service.
- **common**: Shared containers and services that should be running for all services within the project.
   - **containers**: List of shared containers.
   - **docker_compose_file_path**: Docker Compose file for shared services.

## Installation

### Prerequisites

Ensure that you have the following tools installed:

- **Go** (for building the CLI tool)
- **Docker** and **Docker Compose** (for managing containers)

### Steps

1. Clone the main project repository:

   ```bash
   git clone https://github.com/mh-daneshvar/dcli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd dcli
   ```

3. **Tidy up dependencies**:

   ```bash
   go mod tidy
   ```

   This ensures all necessary modules are correctly listed and unused dependencies are removed.

4. Build and install the project using the provided `Makefile`. The following commands are available in the `Makefile`:

   - **Build the project**:

     ```bash
     make build
     ```
     This command builds the binary for your current operating system.

   - **Install the binary**:

     ```bash
     sudo make install
     ```
     This installs the `dcli` binary into `/usr/local/bin` (or another path specified in the `Makefile`).

   - **Build for all supported platforms**:

     ```bash
     make build-all
     ```
     This cross-compiles the binary for Windows, Linux, and macOS across `amd64` architecture. The binaries will be generated in the local directory.

   - **Clean up the built binaries**:

     ```bash
     make clean
     ```
     This removes any previously built binaries from the local directory.

   - **Uninstall the binary**:

     ```bash
     sudo make uninstall
     ```
     This removes the installed binary from the system (default location: `/usr/local/bin`).

   - **Cross-compile for a specific platform** (e.g., Windows `amd64`):

     ```bash
     make build-cross OS=windows ARCH=amd64
     ```
     This cross-compiles the binary for a specific OS and architecture.

5. Clone the helper repository that contains sample projects and services:

   ```bash
   git clone https://github.com/mh-daneshvar/dcli-other-projects.git
   ```

## Usage

### Interactive CLI

Once the `dcli` binary is installed, running the CLI will present you with an interactive menu. This menu is built using the `promptui` library and allows you to choose actions based on available projects and services.

1. Run the CLI tool:

   ```bash
   dcli
   ```

2. The tool will prompt you to select an action. For example:

   ```
   ? Select An Action:
     - Local Development
   ```

3. After selecting **Local Development**, you will be prompted to select a project from the list of projects defined in your YAML file. Example:

   ```
   ? Select the Project:
     - Project One
     - Project Two
   ```

4. Once you select a project, you can choose to start or stop individual services, or stop all services. For example:

   ```
   ? Select the Service:
     - Auth Service
     - Users Service
     - Catalogue Service
     - Stop All Services
   ```

5. Based on your selection:
   - The selected service will either start or stop.
   - If "Stop All Services" is selected, all services related to that project will be stopped.

## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the main repository ([dcli](https://github.com/mh-daneshvar/dcli)).
2. Create a feature branch (`git checkout -b feature/YourFeatureName`).
3. Commit your changes (`git commit -m 'Add YourFeatureName'`).
4. Push to your branch (`git push origin feature/YourFeatureName`).
5. Open a pull request on the main repository.

## License

This project is distributed under the MIT License. See the `LICENSE` file for more details.