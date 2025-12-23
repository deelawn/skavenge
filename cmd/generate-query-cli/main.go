package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

// FunctionInfo holds information about a contract function
type FunctionInfo struct {
	Name       string
	GoName     string
	Parameters []Parameter
	Returns    []Return
}

type Parameter struct {
	Name string
	Type string
}

type Return struct {
	Name string
	Type string
}

type TemplateData struct {
	Functions        []FunctionInfo
	SimpleCommands   []FunctionInfo
	AddressCommands  []FunctionInfo
	TokenCommands    []FunctionInfo
	IndexCommands    []FunctionInfo
	TwoArgCommands   []FunctionInfo
	ThreeArgCommands []FunctionInfo
	SpecialCommands  []FunctionInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <bindings.go> <output.go>\n", os.Args[0])
		os.Exit(1)
	}

	bindingsPath := os.Args[1]
	outputPath := os.Args[2]

	// Parse the bindings file
	functions, err := parseBindings(bindingsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing bindings: %v\n", err)
		os.Exit(1)
	}

	// Categorize functions
	data := categorize(functions)

	// Generate the CLI program
	err = generateCLI(outputPath, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating CLI: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated CLI program: %s\n", outputPath)
}

func parseBindings(path string) ([]FunctionInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var functions []FunctionInfo

	// Find all methods on SkavengeCaller
	ast.Inspect(node, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// Check if this is a method on SkavengeCaller
		if fn.Recv == nil || len(fn.Recv.List) == 0 {
			return true
		}

		recvType := ""
		switch t := fn.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			if ident, ok := t.X.(*ast.Ident); ok {
				recvType = ident.Name
			}
		case *ast.Ident:
			recvType = t.Name
		}

		if recvType != "SkavengeCaller" {
			return true
		}

		// Parse function signature
		funcInfo := FunctionInfo{
			Name:   fn.Name.Name,
			GoName: fn.Name.Name,
		}

		// Parse parameters (skip opts which is always first)
		if fn.Type.Params != nil {
			for i, field := range fn.Type.Params.List {
				if i == 0 {
					// Skip opts parameter
					continue
				}

				typeStr := exprToString(field.Type)
				for _, name := range field.Names {
					funcInfo.Parameters = append(funcInfo.Parameters, Parameter{
						Name: name.Name,
						Type: typeStr,
					})
				}
			}
		}

		// Parse return values
		if fn.Type.Results != nil {
			for _, field := range fn.Type.Results.List {
				typeStr := exprToString(field.Type)
				if len(field.Names) > 0 {
					for _, name := range field.Names {
						funcInfo.Returns = append(funcInfo.Returns, Return{
							Name: name.Name,
							Type: typeStr,
						})
					}
				} else {
					funcInfo.Returns = append(funcInfo.Returns, Return{
						Type: typeStr,
					})
				}
			}
		}

		// Only include functions that return results (not just error)
		if len(funcInfo.Returns) > 1 { // result + error
			functions = append(functions, funcInfo)
		}

		return true
	})

	return functions, nil
}

func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + exprToString(t.Elt)
		}
		return fmt.Sprintf("[%s]%s", exprToString(t.Len), exprToString(t.Elt))
	case *ast.BasicLit:
		return t.Value
	case *ast.StructType:
		return "struct"
	default:
		return fmt.Sprintf("%T", t)
	}
}

func categorize(functions []FunctionInfo) TemplateData {
	data := TemplateData{}

	for _, fn := range functions {
		data.Functions = append(data.Functions, fn)

		switch len(fn.Parameters) {
		case 0:
			data.SimpleCommands = append(data.SimpleCommands, fn)
		case 1:
			paramType := fn.Parameters[0].Type
			if strings.Contains(paramType, "Address") {
				data.AddressCommands = append(data.AddressCommands, fn)
			} else if strings.Contains(paramType, "Int") {
				data.TokenCommands = append(data.TokenCommands, fn)
			} else {
				data.SpecialCommands = append(data.SpecialCommands, fn)
			}
		case 2:
			data.TwoArgCommands = append(data.TwoArgCommands, fn)
		case 3:
			data.ThreeArgCommands = append(data.ThreeArgCommands, fn)
		default:
			data.SpecialCommands = append(data.SpecialCommands, fn)
		}
	}

	// Sort for consistent output
	sort.Slice(data.SimpleCommands, func(i, j int) bool {
		return data.SimpleCommands[i].Name < data.SimpleCommands[j].Name
	})
	sort.Slice(data.AddressCommands, func(i, j int) bool {
		return data.AddressCommands[i].Name < data.AddressCommands[j].Name
	})
	sort.Slice(data.TokenCommands, func(i, j int) bool {
		return data.TokenCommands[i].Name < data.TokenCommands[j].Name
	})
	sort.Slice(data.TwoArgCommands, func(i, j int) bool {
		return data.TwoArgCommands[i].Name < data.TwoArgCommands[j].Name
	})

	return data
}

func generateCLI(outputPath string, data TemplateData) error {
	// Create output directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Execute template
	return cliTemplate.Execute(f, data)
}

var cliTemplate = template.Must(template.New("cli").Funcs(template.FuncMap{
	"lower": strings.ToLower,
	"sub": func(a, b int) int {
		result := a - b
		if result < 0 {
			return 0
		}
		return result
	},
	"camelToKebab": func(s string) string {
		// Convert camelCase to kebab-case for CLI commands
		var result []rune
		for i, r := range s {
			if i > 0 && r >= 'A' && r <= 'Z' {
				result = append(result, '-')
			}
			result = append(result, r)
		}
		return strings.ToLower(string(result))
	},
	"needsBigInt": func(paramType string) bool {
		return strings.Contains(paramType, "big.Int")
	},
	"needsAddress": func(paramType string) bool {
		return strings.Contains(paramType, "Address")
	},
	"needsBytes": func(paramType string) bool {
		return strings.HasPrefix(paramType, "[") && strings.Contains(paramType, "byte")
	},
	"isStruct": func(returnType string) bool {
		return returnType == "struct"
	},
	"formatOutput": func(fn FunctionInfo) string {
		if len(fn.Returns) < 2 {
			return ""
		}
		firstReturn := fn.Returns[0]

		// Check if it's a struct return
		if firstReturn.Type == "struct" {
			return "struct"
		}

		// Check common types
		if strings.Contains(firstReturn.Type, "big.Int") {
			return "bigint"
		}
		if strings.Contains(firstReturn.Type, "Address") {
			return "address"
		}
		if strings.Contains(firstReturn.Type, "string") {
			return "string"
		}
		if strings.Contains(firstReturn.Type, "bool") {
			return "bool"
		}
		if strings.Contains(firstReturn.Type, "byte") {
			return "bytes"
		}
		return "default"
	},
}).Parse(`// Code generated by generate-query-cli. DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/deelawn/skavenge/eth/bindings"
)

const (
	defaultNodeAddress = "http://localhost:8545"
)

func main() {
	// Define flags
	nodeAddress := flag.String("node", defaultNodeAddress, "Ethereum node address")
	flag.Usage = printUsage

	flag.Parse()

	// Get contract address from environment variable
	contractAddressStr := os.Getenv("SKAVENGE_CONTRACT_ADDRESS")
	if contractAddressStr == "" {
		fmt.Fprintf(os.Stderr, "Error: SKAVENGE_CONTRACT_ADDRESS environment variable not set\n")
		os.Exit(1)
	}

	if !common.IsHexAddress(contractAddressStr) {
		fmt.Fprintf(os.Stderr, "Error: Invalid contract address: %s\n", contractAddressStr)
		os.Exit(1)
	}

	contractAddress := common.HexToAddress(contractAddressStr)

	// Get the command
	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	command := args[0]

	// Connect to Ethereum node
	client, err := ethclient.Dial(*nodeAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to node: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// Create contract instance
	contract, err := bindings.NewSkavengeCaller(contractAddress, client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating contract instance: %v\n", err)
		os.Exit(1)
	}

	// Execute command
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	opts := &bind.CallOpts{Context: ctx}

	switch command {
{{- range .SimpleCommands }}
	case "{{ camelToKebab .Name }}":
		exec{{ .Name }}(contract, opts)
{{- end }}
{{- range .AddressCommands }}
	case "{{ camelToKebab .Name }}":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: {{ camelToKebab .Name }} requires an address argument\n")
			os.Exit(1)
		}
		exec{{ .Name }}(contract, opts, args[1])
{{- end }}
{{- range .TokenCommands }}
	case "{{ camelToKebab .Name }}":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: {{ camelToKebab .Name }} requires {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }} argument(s)\n")
			os.Exit(1)
		}
		exec{{ .Name }}(contract, opts, args[1])
{{- end }}
{{- range .TwoArgCommands }}
	case "{{ camelToKebab .Name }}":
		if len(args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: {{ camelToKebab .Name }} requires {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }} arguments\n")
			os.Exit(1)
		}
		exec{{ .Name }}(contract, opts, args[1], args[2])
{{- end }}
{{- range .ThreeArgCommands }}
	case "{{ camelToKebab .Name }}":
		if len(args) < 4 {
			fmt.Fprintf(os.Stderr, "Error: {{ camelToKebab .Name }} requires {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }} arguments\n")
			os.Exit(1)
		}
		exec{{ .Name }}(contract, opts, args[1], args[2], args[3])
{{- end }}
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, ` + "`" + `Skavenge Contract Query Tool

Usage: skavenge-query [options] <command> [arguments]

Environment Variables:
  SKAVENGE_CONTRACT_ADDRESS    Contract address (required)

Options:
  -node string    Ethereum node address (default: http://localhost:8545)

Available Commands:

  No Arguments:
{{- range .SimpleCommands }}
    {{ camelToKebab .Name }}{{ printf "%*s" (sub 40 (len (camelToKebab .Name))) "" }}Get {{ .Name }}
{{- end }}

  Single Argument Commands:
{{- range .AddressCommands }}
    {{ camelToKebab .Name }} <address>{{ printf "%*s" (sub 30 (len (camelToKebab .Name))) "" }}Get {{ .Name }}
{{- end }}
{{- range .TokenCommands }}
    {{ camelToKebab .Name }} <{{ (index .Parameters 0).Name }}>{{ printf "%*s" (sub 30 (len (camelToKebab .Name))) "" }}Get {{ .Name }}
{{- end }}

  Two Argument Commands:
{{- range .TwoArgCommands }}
    {{ camelToKebab .Name }} <{{ range $i, $p := .Parameters }}{{ if $i }}> <{{ end }}{{ $p.Name }}{{ end }}>{{ printf "%*s" (sub 20 (len (camelToKebab .Name))) "" }}Get {{ .Name }}
{{- end }}

Examples:
  export SKAVENGE_CONTRACT_ADDRESS=0x1234567890123456789012345678901234567890
  skavenge-query name
  skavenge-query -node http://localhost:8545 total-supply
  skavenge-query balance-of 0xabcdefabcdefabcdefabcdefabcdefabcdefabcd
  skavenge-query owner-of 1
` + "`" + `)
}

// Command implementations
{{ range .SimpleCommands }}
func exec{{ .Name }}(contract *bindings.SkavengeCaller, opts *bind.CallOpts) {
	result, err := contract.{{ .Name }}(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	{{ $output := formatOutput . -}}
	{{- if eq $output "bigint" }}
	fmt.Println(result.String())
	{{- else if eq $output "address" }}
	fmt.Println(result.Hex())
	{{- else if eq $output "string" }}
	fmt.Println(result)
	{{- else if eq $output "bool" }}
	fmt.Println(result)
	{{- else if eq $output "bytes" }}
	fmt.Printf("0x%x\n", result)
	{{- else }}
	fmt.Printf("%+v\n", result)
	{{- end }}
}
{{ end }}

{{ range .AddressCommands }}
func exec{{ .Name }}(contract *bindings.SkavengeCaller, opts *bind.CallOpts, addressStr string) {
	if !common.IsHexAddress(addressStr) {
		fmt.Fprintf(os.Stderr, "Error: Invalid address: %s\n", addressStr)
		os.Exit(1)
	}
	address := common.HexToAddress(addressStr)
	result, err := contract.{{ .Name }}(opts, address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	{{ $output := formatOutput . -}}
	{{- if eq $output "bigint" }}
	fmt.Println(result.String())
	{{- else if eq $output "address" }}
	fmt.Println(result.Hex())
	{{- else if eq $output "string" }}
	fmt.Println(result)
	{{- else if eq $output "bool" }}
	fmt.Println(result)
	{{- else if eq $output "bytes" }}
	fmt.Printf("0x%x\n", result)
	{{- else }}
	fmt.Printf("%+v\n", result)
	{{- end }}
}
{{ end }}

{{ range .TokenCommands }}
func exec{{ .Name }}(contract *bindings.SkavengeCaller, opts *bind.CallOpts, {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}Str{{ end }} string) {
	{{- range .Parameters }}
	{{- if needsBigInt .Type }}
	{{ .Name }} := new(big.Int)
	{{ .Name }}, ok := {{ .Name }}.SetString({{ .Name }}Str, 10)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ .Name }}: %s\n", {{ .Name }}Str)
		os.Exit(1)
	}
	{{- else if needsBytes .Type }}
	{{ .Name }}Str = strings.TrimPrefix({{ .Name }}Str, "0x")
	var {{ .Name }} {{ .Type }}
	n, err := fmt.Sscanf({{ .Name }}Str, "%x", &{{ .Name }})
	if err != nil || n != 1 {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ .Name }}: %v\n", err)
		os.Exit(1)
	}
	{{- end }}
	{{- end }}
	result, err := contract.{{ .Name }}(opts{{ range .Parameters }}, {{ .Name }}{{ end }})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	{{ $output := formatOutput . -}}
	{{- if eq $output "struct" }}
	// Struct output - print all fields
	fmt.Printf("%+v\n", result)
	{{- else if eq $output "bigint" }}
	fmt.Println(result.String())
	{{- else if eq $output "address" }}
	fmt.Println(result.Hex())
	{{- else if eq $output "string" }}
	fmt.Println(result)
	{{- else if eq $output "bool" }}
	fmt.Println(result)
	{{- else if eq $output "bytes" }}
	fmt.Printf("0x%x\n", result)
	{{- else }}
	fmt.Printf("%+v\n", result)
	{{- end }}
}
{{ end }}

{{ range .TwoArgCommands }}
func exec{{ .Name }}(contract *bindings.SkavengeCaller, opts *bind.CallOpts, {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}Str{{ end }} string) {
	{{- range $idx, $param := .Parameters }}
	{{- if needsBigInt $param.Type }}
	{{ $param.Name }} := new(big.Int)
	{{ $param.Name }}, ok{{ $idx }} := {{ $param.Name }}.SetString({{ $param.Name }}Str, 10)
	if !ok{{ $idx }} {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }}: %s\n", {{ $param.Name }}Str)
		os.Exit(1)
	}
	{{- else if needsAddress $param.Type }}
	if !common.IsHexAddress({{ $param.Name }}Str) {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }} address: %s\n", {{ $param.Name }}Str)
		os.Exit(1)
	}
	{{ $param.Name }} := common.HexToAddress({{ $param.Name }}Str)
	{{- else if needsBytes $param.Type }}
	{{ $param.Name }}Str = strings.TrimPrefix({{ $param.Name }}Str, "0x")
	var {{ $param.Name }} {{ $param.Type }}
	n{{ $idx }}, err{{ $idx }} := fmt.Sscanf({{ $param.Name }}Str, "%x", &{{ $param.Name }})
	if err{{ $idx }} != nil || n{{ $idx }} != 1 {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }}: %v\n", err{{ $idx }})
		os.Exit(1)
	}
	{{- end }}
	{{- end }}
	result, err := contract.{{ .Name }}(opts{{ range .Parameters }}, {{ .Name }}{{ end }})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	{{ $output := formatOutput . -}}
	{{- if eq $output "struct" }}
	// Struct output - print all fields
	fmt.Printf("%+v\n", result)
	{{- else if eq $output "bigint" }}
	fmt.Println(result.String())
	{{- else if eq $output "address" }}
	fmt.Println(result.Hex())
	{{- else if eq $output "string" }}
	fmt.Println(result)
	{{- else if eq $output "bool" }}
	fmt.Println(result)
	{{- else if eq $output "bytes" }}
	fmt.Printf("0x%x\n", result)
	{{- else }}
	fmt.Printf("%+v\n", result)
	{{- end }}
}
{{ end }}

{{ range .ThreeArgCommands }}
func exec{{ .Name }}(contract *bindings.SkavengeCaller, opts *bind.CallOpts, {{ range $i, $p := .Parameters }}{{ if $i }}, {{ end }}{{ $p.Name }}Str{{ end }} string) {
	{{- range $idx, $param := .Parameters }}
	{{- if needsBigInt $param.Type }}
	{{ $param.Name }} := new(big.Int)
	{{ $param.Name }}, ok{{ $idx }} := {{ $param.Name }}.SetString({{ $param.Name }}Str, 10)
	if !ok{{ $idx }} {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }}: %s\n", {{ $param.Name }}Str)
		os.Exit(1)
	}
	{{- else if needsAddress $param.Type }}
	if !common.IsHexAddress({{ $param.Name }}Str) {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }} address: %s\n", {{ $param.Name }}Str)
		os.Exit(1)
	}
	{{ $param.Name }} := common.HexToAddress({{ $param.Name }}Str)
	{{- else if needsBytes $param.Type }}
	{{ $param.Name }}Str = strings.TrimPrefix({{ $param.Name }}Str, "0x")
	var {{ $param.Name }} {{ $param.Type }}
	n{{ $idx }}, err{{ $idx }} := fmt.Sscanf({{ $param.Name }}Str, "%x", &{{ $param.Name }})
	if err{{ $idx }} != nil || n{{ $idx }} != 1 {
		fmt.Fprintf(os.Stderr, "Error: Invalid {{ $param.Name }}: %v\n", err{{ $idx }})
		os.Exit(1)
	}
	{{- end }}
	{{- end }}
	result, err := contract.{{ .Name }}(opts{{ range .Parameters }}, {{ .Name }}{{ end }})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	{{ $output := formatOutput . -}}
	{{- if eq $output "struct" }}
	// Struct output - print all fields
	fmt.Printf("%+v\n", result)
	{{- else if eq $output "bigint" }}
	fmt.Println(result.String())
	{{- else if eq $output "address" }}
	fmt.Println(result.Hex())
	{{- else if eq $output "string" }}
	fmt.Println(result)
	{{- else if eq $output "bool" }}
	fmt.Println(result)
	{{- else if eq $output "bytes" }}
	fmt.Printf("0x%x\n", result)
	{{- else }}
	fmt.Printf("%+v\n", result)
	{{- end }}
}
{{ end }}
`))
