package upgradex

type Options struct {
	Path        string
	WithSQLite  bool
	SkipBuffalo bool
	SkipPop     bool
	SkipApp     bool
	withTests   bool
}
