module {{.PackageName}}

go 1.13

// fix io/fs
replace (
	github.com/spf13/afero => github.com/spf13/afero v1.5.1
	golang.org/x/tools => github.com/golang/tools v0.1.0
)
