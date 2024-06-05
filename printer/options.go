package printer

// Options is the possible options that the
// Printer can have
type Options struct {
	// ShowIcons toggles display of the
	// icons on the end Graph
	ShowIcons bool

	// AlternativeNodeNames will use Node names instead of canonical names
	AlternativeNodeNames bool
}
