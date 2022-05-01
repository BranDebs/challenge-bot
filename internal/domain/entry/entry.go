package entry

// Entry represents the construct used to set conditions as well as update progress.
type Entry struct {
	Name  string `json:"name"`
	Kind  Kind   `json:"kind"`
	Value Value  `json:"value"`
}
