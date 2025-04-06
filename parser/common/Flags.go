package common

type Flags struct {
	Directories string
	Files       string
	Search      string
}

func (flags Flags) DirectoriesPresent() bool {
	return flags.Directories != ""
}

func (flags Flags) FilesPresent() bool {
	return flags.Files != ""
}
