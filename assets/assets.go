package assets

import (
	_ "embed"
)

var (
	//go:embed packed_lemmas.txt
	LemmasData []byte
	//go:embed packed_morphtags.txt
	MorphTagsData []byte
	//go:embed packed_entries.txt
	EntriesData []byte
)
