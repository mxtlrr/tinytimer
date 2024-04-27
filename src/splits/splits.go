/* Splits -- Written by MXTLRR <github.com/mxtlrr>
 * Full logic for reading and parsing splits.
 * Read SPLITS.md for more information on how splits on TinyTimer work */
package splits

import (
	"os"
	"strings"
)

const (
	delim string = ","
)

/* TODO: Implement minutes for those hard events (or if you're
 * a beginner :^) */

type Split_t struct {
	TIME_SECONDS  int    /* Amount of time taken in seconds */
	TIME_MILLISEC int    /* Amount of time taken in milliseconds */
	NAME          string /* Name of split */
}

/* Generate splits from a file.
 * WARNING: OFFSETTED BY 5! MAKE SURE EVERYTHING YOU USE IS n+5! */
func Gen_splits(file string) []Split_t {
	var splits []Split_t

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	// Convert back to string
	better_data := strings.Split(string(data), delim)

	// Now we can bake the array template into a real
	// array, then add a new split to the splits
	splits = make([]Split_t, len(better_data))

	// Add.
	for i := range better_data {
		// Note that TIME_SECONDS and TIME_MILLISEC will both
		// change, once a key is pressed.
		data_split := Split_t{TIME_SECONDS: 0, TIME_MILLISEC: 0, NAME: better_data[i]}
		splits = append(splits, data_split)
	}

	return splits
}
