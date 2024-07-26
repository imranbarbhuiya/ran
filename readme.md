# Ran Programming Language

## Introduction

Ran is a lightweight, expressive programming language designed for simplicity and elegance.

## Why

There's a few reasons why I created this language interpreter

- I wanted to learn Go by building a project
- I was also curious about how programming languages are built
- I wanted to create a simple language that I could use to learn and teach programming concepts
- Telling someone I created a programming language sounds cool 😎

## Features

- **Data Types**: Supports strings, numbers, and booleans and their operations.
- **Variables**: Use `let` to assign values to variables.
- **Control Structures**: Includes `if-else` statements for conditional logic.
- **Mathematical Operations**: Follows proper precedence rules for mathematical operations.
- **Functions**: Define and invoke functions, support higher-order functions, pass functions as parameters, and closures.
- **Collections**: Supports arrays and objects.
- **Macros**: Create macros/pre-processors to generate code.
- **REPL**: Interactive shell for running Ran code.
- **Built-in Functions**: Includes some powerful built-in functions.

## Getting Started

### Hello World

```ran
print("Hello, World!")
```

### Variables and Data Types

```ran
let name = "Ran"
let age = 1
let isActive = true

print(name)
print(age)
print(isActive)
```

### Functions

```ran
let add = fn (a, b) {
    return a + b
}

print(add(2, 3))
```

### Higher-Order Functions

```ran
let applyOperation = fn (a, b, operation) {
    return operation(a, b)
}

let multiply = fn (x, y) {
    return x * y
}

print(applyOperation(3, 4, multiply))
```

### Control Structures

```ran
let num = 10;

if (num > 5) {
    print("Greater than 5")
} else {
    print("Less than or equal to 5")
}
```

### Arrays and Objects

```ran
let fruits = ["Apple", "Banana", "Cherry"]
print(fruits[1])

let person = {
    "name": "Alice",
    "age": 30
}
print(person["name"])
```

### Macros

```ran
let unless = macro(condition, consequence, alternative) {
    quote(if (!(unquote(condition))) {
        unquote(consequence);
    } else {
        unquote(alternative);
    });
};

unless(10 > 5, puts("not greater"), puts("greater"));
```

## What's Next

- [ ] Add line and block comments
- [ ] Add line and column numbers to error messages
- [ ] Add more built-in functions
- [ ] Add more tests
- [ ] A bytecode compiler and a virtual machine
- [ ] A Vs Code extension for syntax highlighting and running Ran code

## Conclusion

I hope you find this project interesting and useful. Feel free to contribute, ask questions, or give feedback. You can also reach out to me on Twitter [@notparbez](https://twitter.com/notparbez). Since I'm new to Go and language design, I'm sure there are many areas where I can improve. I'm excited to learn from you and make this project better together. This isn't a serious project, so I won't implement above features anytime soon.
