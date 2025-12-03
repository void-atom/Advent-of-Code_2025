# Advent of Code 2025 — Go Solutions

This repository contains my solutions for **Advent of Code 2025**, implemented in **Golang**.  
The code uses only the **Go standard library** and follows a modular structure to keep each day's solution clean and reusable.


##  Project Structure
    aoc/
    ├── go.mod
    |── input_files
    |   ├── sample.txt  # Sample data input
    |   ├── day1.txt    # Day 1 input file
    |   ├── day2.txt    # Day 2 input file
    |   └── ...         # Additional days input files
    ├── main.go         # Entry point 
    ├── internal/
    │   └── util.go     # Shared utility functions (e.g., file reading)
    └── days/
        ├── day1.go     # Day 1 solution implementing Day1()
        ├── day2.go     # Day 2 solution implementing Day2()
        └── ...         # Additional days


## Running the Solutions

``` go run main.go ```
