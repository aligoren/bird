package anka

var Static = "static"

func StaticDir(DirName ...string) string {
	if len(DirName) > 0 {
		Static = DirName[0]
	}

	return Static
}

func StaticFiles() string {
	return Static
}