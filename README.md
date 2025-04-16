# VMProtect Wrapper for Go

This package provides a Go wrapper for VMProtect, allowing you to protect your Go applications with VMProtect's anti-tampering and code virtualization features.

## Installation

```bash
go get github.com/MainLibraries/VmProtect-GO
```

## Features

- Virtualization of code sections
- Hardware ID (HWID) generation
- License validation
- SDK function wrapper for Go

## Usage Examples

### Basic Import

```go
import (
    "fmt"
    "github.com/MainLibraries/VmProtect-GO"
)
```

### Getting Hardware ID (HWID)

Hardware IDs are used to identify the machine your software is running on. You can use this for license validation.

```go
package main

import (
    "fmt"
    "github.com/MainLibraries/VmProtect-GO"
)

func main() {
    // Get hardware ID (HWID)
    hwid, err := vmprotect.GetHWID()
    if err != nil {
        fmt.Printf("Error getting HWID: %v\n", err)
        return
    }
    
    fmt.Printf("Hardware ID: %s\n", hwid)
}
```

### Using Begin/End Virtualization

VMProtect allows you to virtualize specific sections of your code to make it more difficult to reverse engineer:

```go
package main

import (
    "fmt"
    "github.com/MainLibraries/VmProtect-GO"
)

func main() {
    fmt.Println("Regular code execution...")
    
    // Start virtualized section
    vmprotect.BeginVirtualization("SecretOperation")
    
    // This code will be protected by VMProtect
    secretValue := calculateSecretValue()
    fmt.Printf("Secret value: %d\n", secretValue)
    
    // End virtualized section
    vmprotect.EndVirtualization()
    
    fmt.Println("Back to regular execution...")
}

func calculateSecretValue() int {
    // This is a sensitive calculation you want to protect
    return 42
}
```

### Using Begin/End Virtualization with Mutation

For even stronger protection, you can use mutation with virtualization:

```go
package main

import (
    "fmt"
    "github.com/MainLibraries/VmProtect-GO"
)

func main() {
    // Start virtualized section with mutation
    vmprotect.BeginVirtualizationWithMutation("LicenseCheck")
    
    // Protected license checking code
    isValid := checkLicense("LICENSE-KEY-123")
    
    // End virtualized section
    vmprotect.EndVirtualization()
    
    if isValid {
        fmt.Println("License is valid!")
    } else {
        fmt.Println("Invalid license!")
    }
}

func checkLicense(key string) bool {
    // Your license validation logic here
    return key == "LICENSE-KEY-123"
}
```

## License Validation Example

```go
package main

import (
    "fmt"
    "github.com/MainLibraries/VmProtect-GO"
)

func main() {
    // Get hardware ID for this machine
    hwid, _ := vmprotect.GetHWID()
    
    // Set the VMProtect serial number (license key)
    err := vmprotect.SetSerialNumber("YOUR-LICENSE-KEY")
    if err != nil {
        fmt.Printf("Error setting serial number: %v\n", err)
        return
    }
    
    // Check if the license is valid for this HWID
    isValid := vmprotect.IsSerialNumberValid()
    
    if isValid {
        fmt.Println("License is valid! Starting application...")
        // Continue with your application logic
    } else {
        fmt.Println("Invalid license! Please purchase a valid license.")
        return
    }
}
```