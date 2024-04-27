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

/* Fixes (#1) :^) */
func insert(a []Split_t, index int, value Split_t) []Split_t {
	// Nil or empty slice or after last element?
	// Just append
	if len(a) == index {
		return append(a, value)
	}

	// Else physically append to value [index]
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

/* TODO: Implement minutes for those hard events (or if you're
 * a beginner :^) */

type Split_t struct {
	TIME_SECONDS  int    /* Amount of time taken in seconds */
	TIME_MILLISEC int    /* Amount of time taken in milliseconds */
	NAME          string /* Name of split */
}

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
	clear(splits) // Clear it up

	// Add.
	for i := range better_data {
		// Note that TIME_SECONDS and TIME_MILLISEC will both
		// change, once a key is pressed.
		data_split := Split_t{TIME_SECONDS: 0, TIME_MILLISEC: 0, NAME: better_data[i]}
		splits = insert(splits, i, data_split)
	}

	// Remove any extraneous
	var current_index int
	for i := range splits {
		if strings.Compare(splits[i].NAME, "") == 0 {
		} else {
			current_index = i
		}
	}

	splits = splits[0 : current_index+1]

	return splits
}
