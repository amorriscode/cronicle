package constants

type Dimensions struct {
	Width  int
	Height int
}

var (
	Sections = []string{"Todo", "Daily", "Brag"}
)

const (
	CONFIG_STORAGE_DIR = "storage_dir"
	CONFIG_USER        = "user"
)

const (
	ERROR_OPEN_FILE   = "Darn, error opening/creating file: %w"
	ERROR_WRITE_FILE  = "Darn, error writing to file: %w"
	ERROR_CLOSE_FILE  = "Darn, error closing file: %w"
	ERROR_LIST_FILE   = "Darn, error listing files: %w"
	ERROR_DELETE_FILE = "Darn, error deleting file: %w"
	ERROR_PROMPT      = "Darn, error with prompt: %w"
)

const MaxLengthDisplayOption = 20
const MaxLengthDetails = 50
