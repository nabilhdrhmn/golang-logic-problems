
# Golang Logic Problems

This repository contains several Go programs that solve different types of problems. Each program is located in its own directory and implements a specific solution as outlined in the problem descriptions.

## Problem Solutions

### 1. String Matching
This program accepts a number of strings and checks for matches between them. It returns the first set of strings that match in a case-insensitive manner.

#### How to Run:
```bash
cd string_matching
go run main.go
```

#### Input Example:
```
4
abcd
acbd
aaab
acbd
```

#### Output Example:
```
2 4
```

### 2. Cashier Change Calculation
This program calculates the change to be given to a customer and breaks it down into the specific denominations of bills and coins. The change is rounded down to the nearest Rp.100.

#### How to Run:
```bash
cd cashier_change
go run main.go
```

#### Input Example:
```
Total belanja seorang customer: Rp 700649
Pembeli membayar: Rp 800000
```

#### Output Example:
```
Kembalian yang harus diberikan kasir: 99351,
dibulatkan menjadi 99300
1 lembar 50000
2 lembar 20000
1 lembar 5000
2 lembar 2000
1 koin 200
1 koin 100
```

### 3. String Validation
This program validates whether a string with special characters `<>{}[]` is well-formed. It checks for properly matching opening and closing brackets.

#### How to Run:
```bash
cd string_validation
go run main.go
```

#### Input Example:
```
{{[<>[{{}}]]}}
```

#### Output Example:
```
TRUE
```

### 4. Leave Eligibility Check
This program checks if an employee is eligible for leave based on the rules regarding company leave policies, including the 180-day rule for new employees.

#### How to Run:
```bash
cd leave_eligibility
go run main.go
```

#### Input Example:
```
Input:
• Jumlah Cuti Bersama = 7
• Tanggal join karyawan = 2021-05-01
• Tanggal rencana cuti = 2021-07-05
• Durasi cuti (hari) = 1
```

#### Output Example:
```
Output:
• False
• Alasan: Karena belum 180 hari sejak tanggal join karyawan
```

## Folder Structure

```
solutions_golang/
│
├── string_matching/
│   └── main.go
├── cashier_change/
│   └── main.go
├── string_validation/
│   └── main.go
└── leave_eligibility/
    └── main.go
```

Each folder contains a `main.go` file, which is the entry point for each respective program.

## How to Run

1. Make sure you have [Go installed](https://golang.org/doc/install).
2. Navigate to the respective folder for the solution you want to run.
3. Use `go run main.go` to execute the program.
