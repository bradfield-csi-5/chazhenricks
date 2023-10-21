package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const src string = `package foo

import (
	"fmt"
	"time"
)

func baz() {
	fmt.Println("Hello, world!")
}

type A int

const b = "testing"

func bar() {
	fmt.Println(time.Now())
}`

// Moves all top-level functions to the end, sorted in alphabetical order.
// The "source file" is given as a string (rather than e.g. a filename).
func SortFunctions(src string) (string, error) {
	f, err := decorator.Parse(src)
	if err != nil {
		panic(err)
	}

	funcDecls, typeDecls, valueDecls, importDecls := extractDecls(f)

	sort.SliceStable(funcDecls, func(i, j int) bool {
		return funcDecls[i].Name.Name < funcDecls[j].Name.Name
	})

	var decls []dst.Decl

	for _, decl := range importDecls {
		decls = append(decls, decl)
	}
	for _, decl := range append(typeDecls, valueDecls...) {
		decls = append(decls, decl)
	}
	for _, decl := range funcDecls {
		decls = append(decls, decl)
	}

	f.Decls = decls

	modifiedString := printFileToString(f)
	fmt.Printf("MODIFIED: %s\n", modifiedString)
	return modifiedString, nil
}

func main() {
	f, err := decorator.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	// Print AST
	err = dst.Fprint(os.Stdout, f, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert AST back to source
	err = decorator.Print(f)
	if err != nil {
		log.Fatal(err)
	}
}

func extractDecls(file *dst.File) ([]*dst.FuncDecl, []*dst.GenDecl, []*dst.GenDecl, []*dst.GenDecl) {
	//create empty struct for function decls
	funcs := make([]*dst.FuncDecl, 0)
	types := make([]*dst.GenDecl, 0)
	values := make([]*dst.GenDecl, 0)
	imports := make([]*dst.GenDecl, 0)
	for _, decl := range file.Decls {
		if fn, ok := decl.(*dst.FuncDecl); ok {
			funcs = append(funcs, fn)
		}
		if genDecl, ok := decl.(*dst.GenDecl); ok {
			//if we have no specs its nothign we can use
			if len(genDecl.Specs) == 0 {
				continue
			}
			switch genDecl.Specs[0].(type) {
			case *dst.TypeSpec:
				types = append(types, genDecl)
			case *dst.ValueSpec:
				values = append(values, genDecl)
			case *dst.ImportSpec:
				imports = append(imports, genDecl)
			}
		}
	}

	return funcs, types, values, imports
}
func printFileToString(file *dst.File) string {
	// Create a new buffer to store the printed code
	var buf strings.Builder

	// Decorator.Fprint function writes the AST to the buffer
	decorator.Fprint(&buf, file)

	// Convert the buffer to a string
	modifiedCode := buf.String()
	return modifiedCode
}
