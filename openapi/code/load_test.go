package code

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	ctx := context.Background()
	batch, err := NewFromFile(ctx, "../testdata/spec.json")
	require.NoError(t, err)

	require.Len(t, batch.Packages(), 1)
	require.Len(t, batch.Services(), 1)
	require.Len(t, batch.Types(), 17)
	commands, ok := batch.packages["commands"]
	require.True(t, ok)

	assert.Equal(t, "commands", commands.FullName())

	ce, ok := commands.services["CommandExecution"]
	require.True(t, ok)
	assert.Equal(t, "commands.CommandExecution", ce.FullName())
	assert.Equal(t, "commandExecution", ce.CamelName())
	assert.Equal(t, "CommandExecution", ce.PascalName())
	assert.Equal(t, commands, ce.Package)

	methods := ce.Methods()
	assert.Equal(t, 7, len(methods))

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
	assert.Equal(t, 17, len(types))

	command := types[2]
	assert.Equal(t, "Command", command.PascalName())
	assert.Equal(t, "command", command.CamelName())
	assert.Equal(t, "commands.Command", command.FullName())
	assert.Equal(t, 4, len(command.Fields()))

	assert.Nil(t, command.Field("nothing"))

	language := command.Field("language")
	assert.NotNil(t, language)
	assert.False(t, language.IsOptionalObject())
	assert.Equal(t, 3, len(language.Entity.Enum()))

	cancel := methods[0]
	cancelResponse := cancel.Response
	assert.Equal(t, "cancel", cancel.Name)
	assert.Equal(t, "cancelResponse", cancelResponse.Name)
	assert.Equal(t, 2, len(cancelResponse.fields))
	assert.Equal(t, "content-type", cancelResponse.fields["content-type"].Name)
	assert.Equal(t, "last-modified", cancelResponse.fields["last-modified"].Name)
}

// This test is used for debugging purposes
func TestMethodsReport(t *testing.T) {
	t.SkipNow()
	ctx := context.Background()
	home, _ := os.UserHomeDir()
	batch, err := NewFromFile(ctx, filepath.Join(home,
		"universe/bazel-bin/openapi/all-internal.json"))
	assert.NoError(t, err)

	for _, pkg := range batch.Packages() {
		for _, svc := range pkg.Services() {
			singleService := strings.EqualFold(pkg.PascalName(), svc.PascalName())
			serviceSingularKebab := svc.Singular().KebabName()
			for _, m := range svc.Methods() {
				var fields []string
				if m.Request != nil {
					for _, f := range m.Request.Fields() {
						flag := fmt.Sprintf("--%s", f.KebabName())
						if f.Entity.IsObject() || f.Entity.MapValue != nil {
							flag = fmt.Sprintf("%%%s", flag)
						}
						fields = append(fields, flag)
					}
				}
				methodWithoutService := (&Named{
					Name: strings.ReplaceAll(
						strings.ReplaceAll(
							m.KebabName(), svc.KebabName(), ""),
						serviceSingularKebab, ""),
				}).KebabName()
				println(fmt.Sprintf("| %s | %s | %s | %s | %v | %s | %s |",
					pkg.KebabName(),
					svc.KebabName(),
					m.KebabName(),
					methodWithoutService,
					singleService,
					m.Operation.Crud,
					strings.Join(fields, ", "),
				))
			}
		}
	}

	assert.Equal(t, len(batch.packages), 1)
}
