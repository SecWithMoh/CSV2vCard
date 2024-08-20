# Go CSV to vCard Converter

This Go application converts a CSV file containing contact information into a vCard (.vcf) format suitable for importing into iPhone contacts. It automatically removes duplicate entries based on phone numbers.

## Features

- **CSV to vCard Conversion:** Converts contact information from CSV to vCard format.
- **Duplicate Removal:** Automatically removes duplicate contacts based on the phone number.
- **Fast and Simple:** Designed for ease of use with minimal setup.

## CSV File Format

The CSV file should follow this structure:

```csv
FirstName,LastName,Company,JobTitle,PhoneNumber
John,Doe,CompanyX,Developer,1234567890
Jane,Smith,CompanyY,Manager,0987654321
```

Make sure that the CSV headers are in English and the fields are ordered exactly as shown.

## Installation

To get started with this project, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/secwithmoh/go-csv-to-vcard.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd go-csv-to-vcard
   ```

3. **Build the project:**

   ```bash
   go build
   ```

This will compile the Go application and generate an executable file in your project directory.

## Usage

1. **Prepare your CSV file:**
   Ensure your CSV file (e.g., `contacts.csv`) is formatted correctly as described above.

2. **Run the application:**

   ```bash
   ./go-csv-to-vcard
   ```

3. **Output:**
   The application will generate a `contacts.vcf` file containing your unique contacts in vCard format, ready for import into iPhone contacts.

### Example

Given a CSV file named `contacts.csv` with the following content:

```csv
FirstName,LastName,Company,JobTitle,PhoneNumber
John,Doe,CompanyX,Developer,1234567890
Jane,Smith,CompanyY,Manager,0987654321
John,Doe,CompanyX,Developer,1234567890
```

The application will produce a `contacts.vcf` file like this:

```vcf
BEGIN:VCARD
VERSION:3.0
N:Doe;John;;;
FN:John Doe
ORG:CompanyX
TITLE:Developer
TEL;TYPE=CELL:1234567890
END:VCARD

BEGIN:VCARD
VERSION:3.0
N:Smith;Jane;;;
FN:Jane Smith
ORG:CompanyY
TITLE:Manager
TEL;TYPE=CELL:0987654321
END:VCARD
```

## Testing

If you'd like to run tests for the project, you can use Go's built-in testing framework. Ensure you have the `testing` package configured in your Go environment, then run:

```bash
go test
```

This will execute any available tests to ensure the code is functioning correctly.

## Compatibility

This project has been tested on:

- **Linux**
- **macOS**
- **Windows**

It should work on any system where Go is supported.

## Contribution

Contributions are welcome! Feel free to open issues, submit pull requests, or suggest new features.

### How to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m "Add some feature"`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a Pull Request.

## License

This project is licensed under the GNU General Public License v3.0. You can find more details in the [LICENSE](LICENSE) file.

### Summary of GPLv3

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see [https://www.gnu.org/licenses/](https://www.gnu.org/licenses/).


## Author

- **SecWithMoh**
- GitHub: [secwithmoh](https://github.com/secwithmoh)
