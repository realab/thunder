package graphql


import (
	"reflect"
)

const (
	SKIP    = "skip"
	INCLUDE = "include"
	IF      = "if"
)

// findDirectiveWithName checks if any of the directives on a field have the sepcified name (eg skip or include)
func findDirectiveWithName(directives []*Directive, name string) *Directive {
	for _, directive := range directives {
		if directive.Name == name {
			return directive
		}
	}
	return nil
}

// shouldIncludeNode validates and checks the value of a skip or include directive
func shouldIncludeNode(directives []*Directive) (bool, error) {
	parseIf := func(d *Directive) (bool, error) {
		args := d.Args.(map[string]interface{})
		if args[IF] == nil {
			return false, NewClientError("required argument in directive not provided: if")
		}

		if _, ok := args[IF].(bool); !ok {
			return false, NewClientError("expected type boolean, found type %v in \"if\" argument", reflect.TypeOf(args["if"]))
		}

		return args[IF].(bool), nil
	}

	skipDirective := findDirectiveWithName(directives, SKIP)
	if skipDirective != nil {
		b, err := parseIf(skipDirective)
		return !b, err
	}

	includeDirective := findDirectiveWithName(directives, INCLUDE)
	if includeDirective != nil {
		return parseIf(includeDirective)
	}

	return true, nil
}
