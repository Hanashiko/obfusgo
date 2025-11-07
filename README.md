# obfusgo

Go code obfuscato.

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

import (
	"encoding/base64"
	"fmt"
)

func d6cabd95b5099cc(a, b int) int {
	return a + b
}

func main() {
	zbde74cab3f950 := "Hello, world!"
	fmt.Println(zbde74cab3f950)
	_ = 7587
	fmt.Println("2+3 =", d6cabd95b5099cc(2, 3))
}
func q70b2cfd9cbeeab(encrypted,

	key string) string {
	F1f38833cebab35, _ :=

		base64.StdEncoding.
			DecodeString(encrypted)
	t5321aed7811f6c, _ := base64.StdEncoding.DecodeString(key)
	for false {
		break
	}
	S923e3389a10bb := make([]byte, len(F1f38833cebab35))
	for false {
		break
	}
	for cdccc2cc8f8cb2c := 0; cdccc2cc8f8cb2c < len(F1f38833cebab35); cdccc2cc8f8cb2c++ {
		S923e3389a10bb[cdccc2cc8f8cb2c] = F1f38833cebab35[cdccc2cc8f8cb2c] ^ t5321aed7811f6c[cdccc2cc8f8cb2c%len(t5321aed7811f6c)]
	}
	for false {
		break
	}
	return string(S923e3389a10bb)
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

Contributions welcome!

## License

MIT License - Use responsibly and legally.

## Disclaimer

Users are responsible for ensuring they have proper authorization before using this tool. The authors assume no liability for misuse.
