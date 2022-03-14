package constants

type Dimensions struct {
	Width  int
	Height int
}

var (
	Sections = []string{"Daily", "Brag"}
)

const (
	STORGAGE_DEFAULT = "storage_dir"
	USER_DEFAULT     = "user"
)

const (
	ERROR_OPEN_FILE  = "Darn, error opening/creating file: %w"
	ERROR_WRITE_FILE = "Darn, error writing to file: %w"
	ERROR_CLOSE_FILE = "Darn, error closing file: %w"
)
