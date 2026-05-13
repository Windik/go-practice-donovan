# Go Practice (Donovan's Book)

This repository contains my solutions to the exercises and code examples from "The Go Programming Language" book by Alan Donovan and Brian Kernighan. I plan to update it as I progress through the book.

## Project Structure

The project is organized by chapters from the book. Each chapter is in its own `chX` directory.

-   `chX/exercises/exerX-Y`: Contains the solution for exercise `Y` from chapter `X`.
-   `chX/theoX`: Contains code examples illustrating theoretical concepts from chapter `X`.

Each directory with a `main.go` file is a separate, runnable program.

## Current Progress

Currently, the repository contains materials from **Chapter 1**.

### Exercises

-   `ch1/exercises/exer1-1`: A program that prints the command name it was invoked with, along with all its arguments.
-   `ch1/exercises/exer1-2`: A program that prints the index and value of each of its command-line arguments, one argument per line.
-   `ch1/exercises/exer1-3`: A program to compare the performance of two ways of string concatenation: using a loop with the `+=` operator and using `strings.Join`.

### Theoretical Examples

-   `ch1/theo1`: "Hello, World".
-   `ch1/theo1-2`: Printing command-line arguments using a loop.
-   `ch1/theo3`: A more efficient way to print command-line arguments using `strings.Join`.

## How to Run

To run any of the programs, navigate to the corresponding directory and execute the `go run .` command.

For example, to run exercise 1-1 from Chapter 1:

```bash
go run ch1/exercises/exer1-1/main.go