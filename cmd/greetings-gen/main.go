package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/tools/go/loader"
)

func loadProgramFromPackage(pkgFullName string) (*loader.Program, error) {
	conf := loader.Config{}
	conf.Import(pkgFullName)
	lprog, err := conf.Load()
	if err != nil {
		return nil, fmt.Errorf("can't load program from package %q: %s",
			pkgFullName, err)
	}

	return lprog, nil
}

func generate(pkgFullName string) error {
	lprog, err := loadProgramFromPackage(pkgFullName)
	if err != nil {
		return err
	}
	pkgInfo := lprog.Package(pkgFullName)
	if pkgInfo == nil {
		return fmt.Errorf("can't get package info from %s", pkgFullName)
	}
	if pkgInfo.Pkg == nil {
		return fmt.Errorf("can't get Pkg from %v", pkgInfo)
	}
	fmt.Println("OK")
	fmt.Println(pkgInfo)
	scope := pkgInfo.Pkg.Scope()
	if scope == nil {
		return fmt.Errorf("can't get scope from %v", pkgInfo.Pkg)
	}
	for _, name := range scope.Names() {
		fmt.Println("-", name)
	}
	return nil
}

func GenerateFromPKG(pkgFullName string) error {
	return generate(pkgFullName)
}

func main() {
	var pkgFullName string

	flag.StringVar(&pkgFullName, "pkg", "", "package full name")

	flag.Parse()

	if err := GenerateFromPKG(pkgFullName); err != nil {
		log.Fatalf("can't generate: %s", err)
	}
}
