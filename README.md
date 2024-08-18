# Go Protobuf UUID Wrapper

This repository provides a Go package that serves as a wrapper for Protocol Buffers (protobuf) with integrated Google UUID support. It is designed for core usage in software systems where UUIDs are essential for unique identification of entities.

## Features

- **Protobuf Integration**: Seamlessly integrates with Protocol Buffers for serialization and deserialization.
- **UUID Generation**: Utilizes Google's `uuid` package to generate and handle UUIDs.
- **Simplicity**: Provides a straightforward interface for generating and embedding UUIDs in protobuf messages.
- **Efficiency**: Designed for high-performance use in production-grade systems.

## Installation

To install the package, use `go get`:

```bash
go get github.com/dmitrorezd/uuid
```

Make sure you have Go installed and set up on your machine.

Usage
Importing the Package

```go
import (
    "github.com/dmitrorezn/uuid"
    "google.golang.org/protobuf/proto"
)
```

Generating UUIDs in Protobuf Messages
Below is a basic example of how to generate and use UUIDs in a protobuf message.

Create uuid directory in vendor pkg and add uuid.proto file from them repository there.
Then Define your protobuf message:

```proto
syntax = "proto3";

import "vendor/uuid/uuid.proto";
package example;

message ExampleMessage {
    UUID id = 1;
    string name = 2;
}
```

Use the Go package to create a new protobuf message with a UUID:


```go
package main

import (
    "fmt"
    "github.com/dmitrorezn/uuid"
    "google.golang.org/protobuf/proto"
)

func main() {
    // Create a new protobuf message
    msg := &example.ExampleMessage{
        Id: uuid.New(),
        Name: "Sample Name",
    }

    // Serialize the message
    data, err := proto.Marshal(msg)
    if err != nil {
        fmt.Println("Error marshaling protobuf message:", err)
        return
    }

    fmt.Println("Serialized message:", data)

    // Deserialize the message
    var newMsg example.ExampleMessage
    err = proto.Unmarshal(data, &newMsg)
    if err != nil {
        fmt.Println("Error unmarshaling protobuf message:", err)
        return
    }

    fmt.Println("Deserialized message:", newMsg.GetId())
    fmt.Println("MessageID:", newMsg.GetId())
    fmt.Println("MessageID String:", newMsg.GetId().String())
    fmt.Println("Message google UUID:", newMsg.GetId().UUID())
}
```

Custom UUID Handling
You can also generate UUIDs manually and assign them to your protobuf message fields:

```go
import "github.com/dmitrorezn/uuid"

func generateCustomUUID() string {
    return uuid.New().String()
}
```

Protobuf UUID Wrapper API
New() string: Generates a new UUID string using the Google UUID package.
Other helper functions for common operations with UUIDs in protobuf messages.
Contributing
We welcome contributions! Please feel free to submit issues, fork the repository, and make pull requests.

Steps to Contribute
Fork the repository.
Create a new branch for your feature or bugfix.
Implement your changes.
Test your changes.
Submit a pull request.
License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Acknowledgements

- **Google Protobuf**
- **Google UUID**

Contact
For any queries, feel free to reach out to us via email at dmitrorezn@gmail.com.