package configuration

type Log struct {
	Filename string

	// megabytes after which new file is created
	MaxSizeMb int

	// number of backups
	MaxBackups int

	MaxAgeDays int
}
