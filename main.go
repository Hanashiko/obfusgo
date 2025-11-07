package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"obfusgo/parser"
	"obfusgo/obfuscation"
)

var (
	inputFile = flag.String("i", "", "Input Go file to obfuscate")
	outputFile = flag.String("o", "", "Output file (default: input_obfuscated.go)")
	methods = flag.String("m", "all", "Obfuscation methods: strings,names,control,dead,all")
	verbose = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	printBanner()

	if *inputFile == "" {
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  obfusgo -i payload.go")
		fmt.Println("  obfusgo -i payload.go -o output.go -m strings,names")
		fmt.Println("  obfusgo -i payload.go -m all -v")
		os.Exit(1)
	}

	if *outputFile == "" {
		if strings.HasSuffix(*inputFile, ".go") && len(*inputFile) > 3 {
			*outputFile = (*inputFile)[:len(*inputFile)-3] + "_obfuscated.go"
		} else {
			*outputFile = *inputFile + "_obfuscated.go"
		}
	}

	if *verbose {
		log.Printf("Input: %s\n", *inputFile)
		log.Printf("Output: %s\n", *outputFile)
		log.Printf("Methods: %s\n", *methods)
	}

	code, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal("Error reading file: %v", err)
	}

	obfuscated, err := obfuscate(code, *methods, *verbose)
	if err != nil {
		log.Fatalf("Obfuscation error: %v", err)
	}

	if err := os.WriteFile(*outputFile, obfuscated, 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	if *verbose {
		log.Println("Obfuscation complete!")
	}

	fmt.Printf("\n[+] Obfuscated file saved: %s\n", *outputFile)
}

func obfuscate(code []byte, methods string, verbose bool) ([]byte, error) {
	obfuscator, err := parser.NewObfuscator(code)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	file := obfuscator.GetAST()

	if strings.TrimSpace(methods) == "all" {
		if verbose {
			log.Println("[*] Applying string encryption...")
		}
		strObf := obfuscation.NewStringObfuscator(verbose)
		strObf.ObfuscateStrings(file)

		if verbose {
			log.Println("[*] Applying name randomization...")
		}

		nameObf := obfuscation.NewNameObfuscator(verbose)
		nameObf.ObfuscateName(file)

		if verbose {
			log.Println("[*] Injecting dead code...")
		}

		deadCode := obfuscation.NewDeadCodeInjector(verbose)
		deadCode.InjectDeadCode(file)

		return obfuscator.Generate()
	}

	methodList := strings.Split(methods, ",")

	for _, method := range methodList {
		method = strings.TrimSpace(method)

		switch method {
		case "strings":
			if verbose {
				log.Println("[*] Applying string encryption...")
			}
			strObf := obfuscation.NewStringObfuscator(verbose)
			strObf.ObfuscateStrings(file)
		case "names":
			if verbose {
				log.Println("[*] Applying name randomization...")
			}
			nameObf := obfuscation.NewNameObfuscator(verbose)
			nameObf.ObfuscateName(file)
		case "dead":
			if verbose {
				log.Println("[*] Injecting dead code...")
			}
			deadCode := obfuscation.NewDeadCodeInjector(verbose)
			deadCode.InjectDeadCode(file)
		default:
		if verbose {
				log.Printf("[!] Unknown method: %s - skipping\n",method)
			}
		}

	}
	return obfuscator.Generate()
}

func printBanner() {
	banner := `
   ___  __    ____                 
  / _ \/ /   / __/_ _____ ____  ___
 / // / _ \ / _/ // (_-</ __/ / _ \
/____/_.__//_/  \_,_/___/\__/  \___/
                                    
obfusgo v0.1 - Go Code Obfuscator
For legal and authorized testing only
`
	fmt.Println(banner)
}
