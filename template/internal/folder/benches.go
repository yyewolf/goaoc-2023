package folder

import (
	"aocli/template/internal/benches"
	"encoding/json"
	"fmt"
	"os"
)

func GetBenches(year string) (out []*benches.BenchOutputGroup) {
	for day := 1; day <= 25; day++ {
		// Search for ./year/day%02d/.bench
		// If not found, continue
		f, err := os.ReadFile(fmt.Sprintf("./%s/day%02d/.bench", year, day))
		if err != nil {
			out = append(out, &benches.BenchOutputGroup{})
			continue
		}

		var parsed benches.BenchOutputGroup

		// Parse bench output
		err = json.Unmarshal(f, &parsed)
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
		}

		out = append(out, &parsed)
	}

	return
}
