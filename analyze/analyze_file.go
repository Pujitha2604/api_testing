package analyze

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

type Endpoint struct {
	Method string
	Path   string
	Result string
}

func analyzeFileForAPIEndpoints(filename string) map[string]Endpoint {
	endpoints := make(map[string]Endpoint)
	fset := token.NewFileSet()

	// Parse the file
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Printf("Error parsing file %s: %v\n", filename, err)
		return endpoints
	}

	// Inspect the AST nodes
	ast.Inspect(node, func(n ast.Node) bool {
		// Check for expression statements
		exprStmt, ok := n.(*ast.ExprStmt)
		if !ok {
			return true
		}

		// Check for call expressions
		callExpr, ok := exprStmt.X.(*ast.CallExpr)
		if !ok {
			return true
		}

		// Check for selector expressions (function calls)
		fun, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || fun.Sel.Name != "Methods" || len(callExpr.Args) == 0 {
			return true
		}

		// Extract method from arguments
		methodLit, ok := callExpr.Args[0].(*ast.BasicLit)
		if !ok {
			return true
		}
		method := strings.Trim(methodLit.Value, "\"")

		// Check for HandleFunc calls
		funSelExpr, ok := fun.X.(*ast.CallExpr)
		if !ok || len(funSelExpr.Args) < 2 {
			return true
		}

		// Check for path literal in HandleFunc arguments
		pathLit, ok := funSelExpr.Args[0].(*ast.BasicLit)
		if !ok {
			return true
		}
		path := strings.Trim(pathLit.Value, "\"")
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}

		// Add endpoint to map
		endpoints[path] = Endpoint{
			Method: method,
			Path:   path,
			Result: "Not Compared",
		}

		return true
	})

	return endpoints
}

func matchEndpoint(handlerEndpoint, newmanEndpoint string) bool {
	handlerBasePath := extractBasePath(handlerEndpoint)
	newmanBasePath := extractBasePath(newmanEndpoint)
	return handlerBasePath == newmanBasePath
}

func extractBasePath(endpoint string) string {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	parts := strings.Split(endpoint, "/")
	if len(parts) > 1 {
		return "/" + parts[1]
	}
	return endpoint
}
