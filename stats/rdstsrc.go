// rdstsrc.go
// R_DST_SRC: The ratio 'number of unique destination addresses' / 'number of unique source addresses'

package stats

import "math"

func init() {
	Register(&RDstSrc{BaseStat{
		name:        "R_DST_SRC",
		description: "Ratio of unique destination addresses to unique source addresses"}})
}

// RDstSrc computes the ratio 'number of unique destination addresses' / 'number of unique source addresses'
type RDstSrc struct {
	BaseStat
}

// Requirement returns the requested counters to compute the stat
func (stat *RDstSrc) Requirement() []string {
	return []string{"NB_UNIQ_DST_ADDR", "NB_UNIQ_SRC_ADDR"}
}

// Compute implements the way to compute the stat from the counters
func (stat *RDstSrc) Compute(ctrvalues []uint64) float64 {
	//ctrvalues[0] -> NB_UNIQ_DST_ADDR
	//ctrvalues[1] -> NB_UNIQ_SRC_ADDR
	if ctrvalues[0] == 0 {
		return 0.
	} else if ctrvalues[1] == 0 {
		return math.NaN()
	} else {
		return float64(ctrvalues[0]) / float64(ctrvalues[1])
	}
}
