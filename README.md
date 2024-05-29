# Random Password Generator

This is a simple command-line tool written in Go to generate a random password of specified length. The generated password includes characters from the following sets:
- Uppercase letters (A-Z)
- Lowercase letters (a-z)
- Numbers (0-9)
- Special characters (!@#$%^&*()-_=+[]{}|;:,.<>?/)

## Installation

Make sure you have Go installed on your system. You can download and install it from [the official Go website](https://golang.org/dl/).

Clone this repository to your local machine:

```bash
git clone https://github.com/Niall1985/PasswordGenGolang.git
```

Navigate to the project directory:

```bash
cd PasswordGenGolang
```

Run the program:

```bash
go run main.go
```

## Usage

By default, the program generates a random password of length 10. You can modify the `len_of_password` variable in the `main` function to change the length of the generated password.

The generated password will be printed to the console.

## License

This project is licensed under the [MIT License](LICENSE).

---

