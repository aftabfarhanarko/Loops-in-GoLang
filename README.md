# Loops in Go (Golang)

In Go, there is only one looping keyword: **`for`**. Go does not have `while` or `do-while` loops. Instead, the `for` statement is flexible enough to handle all looping patterns.

---

## 📌 Table of Contents
- [1. Standard 3-Component Loop](#1-standard-3-component-loop)
- [2. While-Style Loop (Condition Only)](#2-while-style-loop-condition-only)
- [3. Infinite Loop](#3-infinite-loop)
- [4. Loop Control: `break`](#4-loop-control-break)
- [5. Loop Control: `continue`](#5-loop-control-continue)
- [6. Bonus: `for range` Loop](#6-bonus-for-range-loop)
- [How to Run](#how-to-run)

---

## 1. Standard 3-Component Loop
This is the traditional `for` loop consisting of:
- **Init statement**: executed before the first iteration (e.g., `i := 4`)
- **Condition expression**: evaluated before every iteration (e.g., `i <= 100`)
- **Post statement**: executed at the end of every iteration (e.g., `i++`)

```go
for i := 4; i <= 100; i++ {
    fmt.Println(i)
}
```

---

## 2. While-Style Loop (Condition Only)
If you omit the init and post statements, the `for` loop functions like a traditional `while` loop in other languages.

```go
age := 0

for age <= 30 {
    fmt.Println(age)
    age++
}
```

---

## 3. Infinite Loop
Omitting all components creates an infinite loop. It will run endlessly until interrupted or terminated with a `break` or `return`.

```go
for {
    fmt.Println("Running infinitely...")
    // Use break or an exit condition to terminate
    break
}
```

---

## 4. Loop Control: `break`
The `break` statement immediately terminates the execution of the loop and transfers control to the statement following the loop.

```go
for i := 1; i < 8; i++ {
    if i == 6 {
        break // Stops the loop when i reaches 6
    }
    fmt.Println(i)
}
// Output: 1, 2, 3, 4, 5
```

---

## 5. Loop Control: `continue`
The `continue` statement skips the rest of the current iteration and proceeds directly to the next iteration (evaluating the post statement and condition).

```go
for i := 1; i < 8; i++ {
    if i == 6 {
        continue // Skips printing 6 and continues to 7
    }
    fmt.Println(i)
}
// Output: 1, 2, 3, 4, 5, 7
```

---

## 6. Bonus: `for range` Loop
Go also provides the `range` form of the `for` loop to iterate over slices, arrays, maps, strings, or channels.

```go
numbers := []string{"Go", "Python", "Rust"}

for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}
```

---

## 🚀 How to Run

Make sure you have Go installed, then execute:

```bash
go run main.go
```
