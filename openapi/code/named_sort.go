package code

import "sort"

// github.com/databricks/databricks-sdk-go/httpclient/fixtures stub generator uses Named.PascalName() method to come up with
// the best possible field name for generated copy-pastable stubs, though, when this library is attempted to be used together
// with github.com/spf13/viper, we immediately get the following error related to a change in
// golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e as:
// .../entity.go:185:32: type func(a *Field, b *Field) bool of func(a, b *Field) bool {…} does not match inferred type
// ...                        func(a *Field, b *Field) int for func(a E, b E) int,
// because github.com/spf13/viper v0.17+ transitively depends on golang.org/x/exp v0.0.0-20230905200255-921286631fa9

// sorts slice predictably by NamesLikeThis
func pascalNameSort[E interface{ PascalName() string }](things []E) {
	sort.Slice(things, func(i, j int) bool {
		return things[i].PascalName() < things[j].PascalName()
	})
}

// sorts slice predictably by package_names_and.ClassNamesLikeThis
func fullNameSort[E interface{ FullName() string }](things []E) {
	sort.Slice(things, func(i, j int) bool {
		return things[i].FullName() < things[j].FullName()
	})
}
