# obfusgo

Go code obfuscator.

⚠️ **Legal Notice**: This tool is intended for legal competitions and authorized security testing only. Always obtain proper authorization before use.

## Features

- **String Encryption** - XOR encrypts all string literals with random keys
- **Name Randomization** - Randomizes variable, function, and type names
- **Dead Code Injection** - Injects unreachable code to change binary signatures
- **CLI Interface** - Simple command-line usage

## Project Structure

```
obfusgo/
├── main.go
├── parser/
│   └── parser.go          # AST parsing and generation
└── obfuscation/
    ├── strings.go         # String encryption
    ├── names.go           # Name randomization
    └── deadcode.go        # Dead code injection
```

## Installation

```bash
git clone https://github.com/Hanashiko/obfusgo
cd obfusgo
go build
```

## Usage

Basic usage:
```bash
./obfusgo -i payload.go
```

Specify output file:
```bash
./obfusgo -i payload.go -o obfuscated.go
```

Select specific methods:
```bash
./obfusgo -i payload.go -m strings,names
```

Verbose mode:
```bash
./obfusgo -i payload.go -m all -v
```

## Obfuscation Methods

- `strings` - Encrypts all string literals
- `names` - Randomizes identifiers
- `dead` - Injects dead code
- `all` - Apply all methods (default)

## Example

**Before:**
```go
package main

import "fmt"

func main() {
    message := "Hello CTF!"
    fmt.Println(message)
}
```

**After:**
```go
package main

import "fmt"

func decrypt(encrypted, key string) string {
    // Decryption logic...
}

func A7f3e9d2a1b4c5() {
    B8e4f1a2d3c6 := decrypt("kR3pL...", "9xT2m...")
    if false { _ = 0 }
    fmt.Println(B8e4f1a2d3c6)
}
```

## How It Works

1. **AST Parsing** - Parses Go source into Abstract Syntax Tree
2. **Transformation** - Applies selected obfuscation techniques
3. **Code Generation** - Generates obfuscated Go code

### String Encryption

- Generates random XOR key per string
- Base64 encodes encrypted data
- Injects decrypt function
- Helps bypass signature-based detection

### Name Randomization

- Maps original names to random identifiers
- Preserves Go keywords and exported names
- Makes reverse engineering harder

### Dead Code Injection

- Adds unreachable code blocks
- Changes binary signature
- Doesn't affect execution

## Limitations

- May not work with reflection-heavy code
- Generated code might be harder to debug
- Some optimizations may remove dead code

## Future Enhancements

- [ ] Control flow flattening
- [ ] Import obfuscation
- [ ] Anti-debug techniques
- [ ] Polymorphic code generation
- [ ] Automatic decrypt function injection
- [ ] Support for entire packages

## Contributing

This is an educational project for CTF competitions. Contributions welcome!

## License

MIT License - Use responsibly and legally.

## Disclaimer

This tool is for authorized security testing and CTF competitions only. Users are responsible for ensuring they have proper authorization before using this tool. The authors assume no liability for misuse.
