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
type Response struct {
	Code int `json:"code"`
}

func analyzeFileForAPIEndpoints(filename string) map[string]Endpoint {
	endpoints := make(map[string]Endpoint)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Printf("Error parsing file %s: %v\n", filename, err)
		return endpoints
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if exprStmt, ok := n.(*ast.ExprStmt); ok {
			if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
				if fun, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
					if fun.Sel.Name == "Methods" && len(callExpr.Args) > 0 {
						if methodLit, ok := callExpr.Args[0].(*ast.BasicLit); ok {
							method := strings.Trim(methodLit.Value, "\"")
							if funSelExpr, ok := fun.X.(*ast.CallExpr); ok {
								if funHandleFunc, ok := funSelExpr.Fun.(*ast.SelectorExpr); ok {
									if funHandleFunc.Sel.Name == "HandleFunc" && len(funSelExpr.Args) >= 2 {
										if pathLit, ok := funSelExpr.Args[0].(*ast.BasicLit); ok {
											path := strings.Trim(pathLit.Value, "\"")
											if !strings.HasPrefix(path, "/") {
												path = "/" + path
											}
											endpoints[path] = Endpoint{
												Method: method,
												Path:   path,
												Result: "Not Compared",
											}
										}
									}
								}
							}
						}
					}
				}
			}
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
	parts := strings.Split(endpoint, "/")
	if len(parts) > 1 {
		return "/" + parts[1]
	}
	return endpoint
}
