package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/samber/lo"
)

func GetOptionAliasMap(args []string) (map[string]string, error) {
	// command-specific options
	// HACK: append `-h` to get help message of specified command
	specOpts, err := getOptionAliasMap(append(args, "-h"))
	if err != nil {
		return nil, fmt.Errorf("failed to get command-specific options: %w", err)
	}

	// global options
	globalOpts, err := getOptionAliasMap([]string{"options"})
	if err != nil {
		return nil, fmt.Errorf("failed to get global options: %w", err)
	}

	return lo.Assign(globalOpts, specOpts), nil
}

func getOptionAliasMap(args []string) (map[string]string, error) {
	cmd := exec.Command("kubectl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to get option list: args: `%s`: %w\n\n%s", strings.Join(args, " "), err, out)
	}

	optLines := trimOptionText(string(out))

	return parseOptionAliasMap(optLines), nil
}

func trimOptionText(out string) []string {
	lines := lo.Map(strings.Split(out, "\n"), func(line string, _ int) string {
		return strings.TrimSpace(line)
	})

	// extract lines that start with `-` or `--`
	// ex: `-o, --output='':`
	return lo.Filter(lines, func(line string, _ int) bool {
		return strings.HasPrefix(line, "-")
	})
}

func parseOptionAliasMap(lines []string) map[string]string {
	aliases := map[string]string{}

	for _, line := range lines {
		if alias, option, ok := parseOptionAlias(line); ok {
			aliases[option] = alias
		}
	}

	return aliases
}

func parseOptionAlias(line string) (alias string, option string, ok bool) {
	if !strings.Contains(line, ",") {
		return "", "", false
	}

	// trim default value (ex: `=[]:` in `-f, --filename=[]:`)
	optionsText := strings.Split(line, "=")[0]

	// split options
	options := strings.Split(optionsText, ", ")
	if len(options) != 2 {
		return "", "", false
	}

	return options[0], options[1], true
}
