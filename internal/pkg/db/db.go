package db

type Db struct {
	name   string
	dsn    string
	dir    string
	tables map[string]Table
}

type Table struct {
	name  string
	idc   string
	tc    string
	fname string
}
