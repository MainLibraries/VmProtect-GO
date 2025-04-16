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
    hwid := vmprotect.GetCurrentHWID()
    // Note: GetCurrentHWID does not return an error.
    
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
    vmprotect.End()
    
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
    vmprotect.BeginMutation("LicenseCheck")
    
    // Protected license checking code
    isValid := checkLicense("LICENSE-KEY-123")
    
    // End virtualized section
    vmprotect.End()
    
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
    hwid := vmprotect.GetCurrentHWID()
    fmt.Printf("Current HWID: %s\n", hwid) // Optional: Print HWID for debugging
    
    // Set the VMProtect serial number (license key)
    // SetSerialNumber returns an int status, 0 usually means success.
    status := vmprotect.SetSerialNumber("YOUR-LICENSE-KEY")
    if status != 0 { // Or check against specific error codes if available
        fmt.Printf("Error setting serial number, status: %d\n", status)
        return
    }
    
    // Check if the license is valid
    state := vmprotect.GetSerialNumberState()
    
    if state == vmprotect.SerialStateSuccess {
        fmt.Println("License is valid! Starting application...")
        // You can also retrieve detailed serial number data if needed:
        // data, ok := vmprotect.GetSerialNumberData()
        // if ok {
        //     fmt.Printf("License User: %s, Expires: %d-%d-%d\n", data.UserName, data.ExpireDate.Year, data.ExpireDate.Month, data.ExpireDate.Day)
        // }
        // Continue with your application logic
    } else {
        fmt.Printf("Invalid license! State: %d. Please purchase a valid license.\n", state)
        // You might want to print more specific error messages based on the state flags, e.g.:
        // if state&vmprotect.SerialStateFlagDateExpired != 0 { fmt.Println("License has expired.") }
        // if state&vmprotect.SerialStateFlagBadHWID != 0 { fmt.Println("License is for a different hardware.") }
        return
    }
}
```
