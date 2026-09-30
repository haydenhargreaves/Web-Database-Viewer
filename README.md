# Web-Database-Viewer

## Table of Contents

- [Introduction](#introduction)
- [Version](#version)
- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Docker](#docker)
    - [Supported Drivers](#supported-drivers)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This is a simple web database viewer that allows you to view and edit a database in a web browser. Connections
can be made to the database using the convenient connection creation interface, and later modified and deleted 
with the manage connections interface.

Connections are stored in the web browser's local storage and are not stored on the server or any external database.
For extra security you may delete the database connection information from the browser's local storage upon each use.
Otherwise, leaving them in the local storage will allow you to quickly connect to the database without having to re-enter
the details each time.

When a query is ran on the database selected, the connection is created and the query is ran. The connection is then closed
and the results are displayed in the web browser. This is done to prevent any security risks that may arise from leaving
a connection open to the database.

## Version

This is version 1.0 of the web database viewer. All endpoints will be proceded with the version number in the URL.

e.g. `http(s)://hostname:port/v1/...`

**Versions Include:**
- v1

## Features

- Connect to many databases *(Supported database drivers can be seen below)*
- View tree structure of database
- View table restraints such as primary keys, foreign keys, types, etc.
- View enum names and values
- Quickly query tables with the quick query buttons on each table
- Quickly select columns from tables using the tree structure
- Toggle between connections with the click of a button
- View references to other tables based on foreign keys in tree table
- Automatically run SQL queries when input box is changed or manually run them yourself
- View results of SQL queries in a table format

## Getting Started

This project was built in [GoLang](https://go.dev) using the [Gin web framework](https://github.com/gin-gonic/gin). The front-end was built using HTML, 
with interaction powered by [HTMX](https://htmx.org), and styled with [Tailwind CSS](https://tailwindcss.com).

### Prerequisites

If you wish to make changes to this project, you will need the following installed on your machine:

- [GoLang](https://go.dev) Version 1.16 or higher (built with 1.22.5)
- [Tailwind CSS](https://tailwindcss.com) (for styling)

If you wish to run this project locally, without any changes, you will not need anything other than 
the project files which can be found in the [releases page](https://github.com/Azpect3120/Web-Database-Viewer/releases).

### Installation

To make changes to this project you will need to clone the repository and run the following command:

```bash
go mod tidy
```

A Tailwind executable and various scripts are included in the project to make building the project easier.
To compile the script for final use, run the following command:

```bash
./tools/styles/compile.sh
```

To start the Tailwind CSS watcher for live reloading of styles, simply run the following command:

```bash
./tools/styles/watch.sh
```

Finally, to run the project, simply run the following command:

```bash
go run cmd/web_server.go
```

You can then see a live running version of the application running on your localhost on port 3001. The port 
can be changed in the `cmd/web_server.go` file. For now there is no requirement for any environment variables
or setup. This may be added in the future, but for now, the project is simple enough to run without any setup 
(assuming port 3001 is available).

### Docker

Published images are available from the GitHub Container Registry:

```bash
docker run --rm -p 3001:3001 ghcr.io/haydenhargreaves/web-database-viewer:latest
```

The application is available at `http://localhost:3001/v1/web/view`. Its health endpoint is
`http://localhost:3001/v1/api/status`.

Pull requests build and smoke-test the image without publishing it. Pushes to `master` publish `latest` and
an immutable commit-SHA tag. Semantic version tags, for example `v1.0.0`, additionally publish version tags.


### Supported Drivers

Since this project was built to support multiple database drivers, the connection details, query information, and
tree display are all based on the driver used. Once a connection is made, the query input table will always work
but for drivers that are not supported the table/enum trees will not be displayed. Fully supported drivers will
have all features available to them. Partially supported drivers will have some features available to them, but not
the full range of features.

**Fully Supported Drivers:**
- PostgreSQL
- MySQL
- MariaDB
- SQLite3

**Partially Supported Drivers:**
- SQL Server
- Oracle
- DB2

## Contributing

Contributions are always welcome! Please make sure to include a comment in your PR so I know what the 
purpose of the change is. If you'd like to contribute to this project, please follow these steps:

1. Fork the project.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Test your changes thoroughly.
5. Create a pull request.

## License

The project is licensed under Azpect3120 the **MIT License**
