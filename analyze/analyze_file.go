package analyze

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Endpoint struct {
	Method string
	Path   string
	Result string
}

type NewmanReport struct {
	Run Run `json:"run"`
}

type Run struct {
	Executions []Execution `json:"executions"`
}

type Execution struct {
	Item     Item     `json:"item"`
	Response Response `json:"response"`
}

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method string `json:"method"`
	URL    URL    `json:"url"`
}

type URL struct {
	Path []string `json:"path"`
}

type Response struct {
	Code int `json:"code"`
}

func analyzeFileForAPIEndpoints(rootDir string) map[string]Endpoint {
	endpoints := make(map[string]Endpoint)

	processFile := func(filename string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(filename, ".go") {

			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
			if err != nil {
				log.Printf("Error parsing file %s: %v\n", filename, err)
				return nil
			}

			for _, decl := range node.Decls {
				if fdecl, ok := decl.(*ast.FuncDecl); ok {
					if fdecl.Doc != nil {
						var method, path string
						for _, comment := range fdecl.Doc.List {
							text := strings.TrimSpace(comment.Text)
							if strings.HasPrefix(text, "//@Method:") {
								method = strings.TrimSpace(strings.TrimPrefix(text, "//@Method:"))
							} else if strings.HasPrefix(text, "//@Route:") {
								path = strings.TrimSpace(strings.TrimPrefix(text, "//@Route:"))
							}
						}
						if method != "" && path != "" {
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
		return err
	}

	err := filepath.Walk(rootDir, processFile)
	if err != nil {
		log.Printf("Error walking directory %s: %v\n", rootDir, err)
	}

	return endpoints
}

func matchEndpoint(handlerEndpoint, newmanEndpoint string) bool {
	pattern := "^" + regexp.QuoteMeta(handlerEndpoint) + "$"
	matched, _ := regexp.MatchString(pattern, newmanEndpoint)
	return matched
}
