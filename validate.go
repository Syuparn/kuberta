package main

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateResourceNames(args []string, aliases map[string]string) error {
	expectedArgs := make([]string, len(args))
	copy(expectedArgs, args)

	for i, arg := range args {
		if alias, ok := aliases[arg]; ok {
			expectedArgs[i] = alias
		}

		// detect resource prefix like `service/foo`
		prefix := strings.Split(arg, "/")[0]
		if alias, ok := aliases[prefix]; ok {
			expectedArgs[i] = strings.Replace(arg, prefix, alias, 1)
		}
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		return fmt.Errorf("too long! should be `kubectl %s`", strings.Join(expectedArgs, " "))
	}

	return nil
}
