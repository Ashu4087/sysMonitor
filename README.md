# sysMonitor

A fast and efficient command-line interface (CLI) system monitor written in Go. This tool provides real-time insights into system health directly from the terminal, bypassing the need for heavy GUI applications.

## Core Features

- Real-time RAM, CPU, and Disk usage tracking.
- Seamless terminal UI with continuous loop updates.
- Concurrent metrics fetching using Go routines for high performance.

Built leveraging Go's robust standard library and system utilities.

## Prerequisites

- Go 1.20 or higher
- `make` utility (for running build scripts)

## How to Run & Build

This project uses a `Makefile` for easy command execution.

### For Development
To run the application directly without building a binary:
```bash
make run
```

### To Build the Executable
To compile the project into an executable binary:
```bash
make build
```
*The generated binary will be placed in the `bin/` directory.*

### To Clean the Build
To remove the generated binaries and clean the directory:
```bash
make clean
```