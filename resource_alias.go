package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func GetResourceAliasMap() (map[string]string, error) {
	cmd := exec.Command("kubectl", "api-resources", "--no-headers")
	b, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get resource aliases: %w", err)
	}

	records := splitOutput(string(b))

	return parseAliasMap(records)
}

func splitOutput(b string) [][]string {
	spaces := regexp.MustCompile(`\s+`)

	lines := strings.Split(b, "\n")
	return lo.Map(lines, func(line string, _ int) []string {
		return spaces.Split(line, -1)
	})
}

func parseAliasMap(records [][]string) (map[string]string, error) {
	m := map[string]string{}

	for _, record := range records {

		if len(record) < 5 {
			// no shortNames defined
			continue
		}

		shortNames := strings.Split(record[1], ",")
		name := record[0]
		kind := record[4]

		for _, shortName := range shortNames {
			// plural
			m[name] = shortName

			// singular
			m[strings.ToLower(kind)] = shortName
		}
	}

	return m, nil
}
