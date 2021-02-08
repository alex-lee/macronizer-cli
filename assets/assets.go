package assets

import (
	_ "embed"
)

var (
	//go:embed packed_lemmas.txt.gz
	LemmasData []byte
	//go:embed packed_morphtags.txt.gz
	MorphTagsData []byte
	//go:embed packed_entries.txt.gz
	EntriesData []byte
)
