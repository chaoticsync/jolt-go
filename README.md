# Jolt-Go

Jolt-Go is a Go implementation of the Jolt JSON transformation library, which provides a simple and powerful way to transform JSON data.

## Features

- Shift transformation: Move data from one location to another
- Default transformation: Set default values for missing fields
- Remove transformation: Remove specified fields from the output
- Cardinality transformation: Control array cardinality (ONE/MANY)
- Sort transformation: Sort arrays based on their values
- Chain transformation: Combine multiple transformations

## Installation

```bash
# Install specific version
go get github.com/chaoticsync/jolt-go@v0.1.0

# Install latest version
go get github.com/chaoticsync/jolt-go
```

## Usage

### Command Line

```bash
# If you have the source code:
go run cmd/main.go <input-file> <spec-file>

# If installed as a package:
jolt-go <input-file> <spec-file>
```

The program takes two arguments:
- `input-file`: Path to the JSON file containing the input data
- `spec-file`: Path to the JSON file containing the JOLT transformation specification

### Programmatic Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/chaoticsync/jolt-go/pkg/transformer"
)

func main() {
    // Define input JSON
    input := map[string]interface{}{
        "name": "John",
        "age":  30,
    }

    // Define shift specification
    shiftSpec := &transformer.Spec{
        Operation: transformer.Shift,
        Spec: map[string]interface{}{
            "name": "person.name",
            "age":  "person.age",
        },
    }

    // Create transform
    transform, err := transformer.NewTransform(shiftSpec)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Apply transformation
    result, err := transform.Apply(input)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Print result
    output, _ := json.MarshalIndent(result, "", "  ")
    fmt.Println(string(output))
}
```

## Transformation Types

### Shift

The shift transformation moves data from one location to another in the JSON structure.

Example specification:
```go
spec := &transformer.Spec{
    Operation: transformer.Shift,
    Spec: map[string]interface{}{
        "name": "person.name",
        "age":  "person.age"
    },
}
```

### Default

The default transformation sets default values for missing fields.

Example specification:
```go
spec := &transformer.Spec{
    Operation: transformer.Default,
    Spec: map[string]interface{}{
        "status": "active",
        "role": "user"
    },
}
```

### Remove

The remove transformation removes specified fields from the output.

Example specification:
```go
spec := &transformer.Spec{
    Operation: transformer.Remove,
    Spec: map[string]interface{}{
        "password": nil,
        "secret": nil
    },
}
```

### Cardinality

The cardinality transformation controls array cardinality.

Example specification:
```go
spec := &transformer.Spec{
    Operation: transformer.Cardinality,
    Spec: map[string]interface{}{
        "items": "ONE"
    },
}
```

### Sort

The sort transformation sorts arrays based on their values.

Example specification:
```go
spec := &transformer.Spec{
    Operation: transformer.Sort,
    Spec: map[string]interface{}{
        "items": {}
    },
}
```

## License

MIT