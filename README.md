# Go-Get-My-Password

=======================

A simple password manager that runs on a looping Command-Line Interface (CLI) built with Go.

## Overview

---

This password manager allows users to store, retrieve, and manage their passwords in a secure and convenient way. The CLI provides a menu-driven interface for users to interact with the password manager.

## Features

---

- Store passwords securely using AES encryption
- Retrieve passwords by domain name/URI
- Add, edit, and delete password entries
- Generate strong, unique passwords (coming soon)

## Usage

---

1. Run the CLI using `go run main.go`
2. Follow the menu prompts to perform desired actions

## Requirements

---

- Go 1.13 or later
- The following dependencies:
  - `github.com/mattn/go-sqlite3` for SQLite database operations
  - `golang.org/x/crypto/pbkdf2` for password-based key derivation

## Security

---

- Passwords are stored securely using AES encryption
- Users are prompted to enter a master password to access the password manager

## Contributing

---

Contributions are welcome! Please submit a pull request with your changes and a brief description of what you've added or fixed.

## License

---
