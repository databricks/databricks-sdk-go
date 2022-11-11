package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json", []string{})
	require.NoError(t, err)

	require.Len(t, batch.Pkgs(), 1)
	require.Len(t, batch.Services(), 1)
	require.Len(t, batch.Types(), 14)
	commands, ok := batch.Packages["commands"]
	require.True(t, ok)

	assert.Equal(t, "commands", commands.FullName())

	ce, ok := commands.services["CommandExecution"]
	require.True(t, ok)
	assert.Equal(t, "commands.CommandExecution", ce.FullName())
	assert.Equal(t, "commandExecution", ce.CamelName())
	assert.Equal(t, "CommandExecution", ce.PascalName())
	assert.Equal(t, commands, ce.Package)

	methods := ce.Methods()
	assert.Equal(t, 6, len(methods))

	execute := methods[5]
	assert.Equal(t, "execute", execute.Name)
	assert.Equal(t, "Execute", execute.PascalName())
	assert.Equal(t, "Post", execute.TitleVerb())
	assert.Nil(t, execute.Shortcut())

	wait := execute.Wait()
	require.NotNil(t, wait)
	binding := wait.Binding()
	assert.Equal(t, 3, len(binding))

	assert.Equal(t, "finished", wait.Success()[0].CamelName())
	assert.Equal(t, "CommandStatus", wait.Failure()[0].Entity.PascalName())
	assert.Equal(t, 0, len(wait.MessagePath()))

	types := commands.Types()
	assert.Equal(t, 14, len(types))

	command := types[1]
	assert.Equal(t, "Command", command.PascalName())
	assert.Equal(t, "command", command.CamelName())
	assert.Equal(t, "commands.Command", command.FullName())
	assert.Equal(t, 4, len(command.Fields()))

	assert.Nil(t, command.Field("nothing"))

	language := command.Field("language")
	assert.NotNil(t, language)
	assert.False(t, language.IsOptionalObject())
	assert.Equal(t, 3, len(language.Entity.Enum()))
}

// This test is used for debugging purposes
func TestBasicDebug(t *testing.T) {
	t.SkipNow()
	home, _ := os.UserHomeDir()
	batch, err := NewFromFile(filepath.Join(home,
		"universe/bazel-bin/openapi/all-internal.json"), []string{})
	assert.NoError(t, err)

	for _, pkg := range batch.Pkgs() {
		if pkg.Name != "workspaceconf" {
			continue
		}
		pkg.Types()
		for _, svc := range pkg.Services() {
			for _, m := range svc.Methods() {
				pg := m.Pagination()
				if pg != nil {
					t.Logf("pg entity: %s", pg.Entity.PascalName())
				}
				w := m.Wait()
				if w != nil {
					t.Logf("wait: %v", w.Success())
				}
			}
		}
	}

	assert.Len(t, batch.Packages, 1)
}
