# Project 01 --- Student Grade Management System (CLI)

## Overview

This project is a command-line application written in Go that simulates
a simple student grade management system.

The application was designed as a beginner backend-oriented exercise to
practice Go fundamentals such as structs, slices, functions, loops,
input handling, sorting, and basic CRUD operations---all without using a
database or external libraries.

## Scenario

A university lecturer wants a lightweight application to manage students
and their grades.

Since the number of students is relatively small, there is no need for a
database yet. All data is stored in memory while the application is
running.

The goal is to provide a simple interface that allows the lecturer to:

-   Add new students
-   View all students
-   Search for students
-   Calculate averages
-   View overall statistics

## Functional Requirements

### Main Menu

``` text
=============================
 Student Grade Management
=============================

1. Add Student
2. View Students
3. Search Student
4. Show Statistics
5. Exit
```

The application keeps running until the user selects **Exit**.

### Student Information

Each student contains:

-   Student Name
-   Student ID
-   List of Grades

### Feature 1 --- Add Student

The user can enter:

-   Student Name
-   Student ID
-   Number of grades
-   Individual grades

### Feature 2 --- View Students

Display every student's:

-   Name
-   Student ID
-   Grades
-   Average grade
-   Pass/Fail status

Passing rule:

-   Average ≥ 60 → PASS
-   Average \< 60 → FAIL

### Feature 3 --- Search Student

Search students using their Student ID.

If found, display the complete student information.

If not found:

``` text
Student not found.
```

### Feature 4 --- Statistics

Display:

-   Total number of students
-   Student with highest average
-   Student with lowest average
-   Overall average of all students

## Technical Constraints

-   Standard Go library only
-   No database
-   No JSON persistence
-   No web framework
-   No external packages
-   Data stored in memory using slices

## Suggested Data Structure

``` go
type Student struct {
    Name   string
    ID     string
    Grades []float64
}

var students []Student
```

## Learning Objectives

This project was designed to practice:

-   Go project structure
-   Packages and modules
-   Structs
-   Slices
-   Functions
-   Loops
-   Conditional statements
-   User input
-   Searching algorithms
-   Sorting collections
-   Basic statistics
-   Code organization across multiple files

## Bonus Challenges

-   Prevent duplicate Student IDs
-   Edit existing student grades
-   Delete students
-   Sort students by average grade (highest to lowest)
-   Persist data using JSON files

## Skills Demonstrated

-   CRUD operations (Create, Read, Update, Delete)
-   Basic data modeling
-   Business logic implementation
-   Command-line application development
-   Problem solving and debugging in Go

## Author's Note

This is the first project in my Go learning journey. The goal wasn't to
build a production-ready application, but to strengthen my understanding
of Go fundamentals before moving into backend web development. Future
projects will gradually introduce file persistence, REST APIs,
databases, authentication, and deployment.
