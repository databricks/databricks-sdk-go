package workspace

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/commands"
)

var b64 = base64.StdEncoding

// PythonNotebookOverwrite crafts Python import notebook request
// also by trimming the code specified in the second argument
func PythonNotebookOverwrite(path, content string) Import {
	content = commands.TrimLeadingWhitespace(content)
	request, _ := PythonNotebookOverwriteReader(path,
		strings.NewReader(content))
	return request
}

func PythonNotebookOverwriteReader(path string, r io.Reader) (Import, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return Import{}, fmt.Errorf("read: %w", err)
	}
	return Import{
		Path:      path,
		Overwrite: true,
		Format:    ExportFormatSource,
		Language:  LanguagePython,
		Content:   b64.EncodeToString(raw),
	}, nil
}

func (r *ExportResponse) Bytes() ([]byte, error) {
	return b64.DecodeString(r.Content)
}
