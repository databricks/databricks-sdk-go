// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

// all definitions in this file are in alphabetical order

type CreateScriptResponse struct {
	ScriptId any /* MISSING TYPE */ `json:"script_id,omitempty"`
}

type DeleteScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" path:"script_id"`
}

type GetScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" path:"script_id"`
}

type GlobalInitScriptCreateRequest struct {
	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name"`

	Position int `json:"position,omitempty"`

	Script string `json:"script"`
}

type GlobalInitScriptDetails struct {
	// When the script was created, as a Unix timestamp in milliseconds.
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	CreatedBy string `json:"created_by,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`

	Position int `json:"position,omitempty"`

	ScriptId string `json:"script_id,omitempty"`
	// When the script was updated, as a Unix timestamp in milliseconds.
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	UpdatedBy string `json:"updated_by,omitempty"`
}

type GlobalInitScriptUpdateRequest struct {
	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`

	Position int `json:"position,omitempty"`

	Script string `json:"script,omitempty"`
	// The ID of the global init script.
	ScriptId string `json:"-" path:"script_id"`
}

// The Base64-encoded content of the script.

// Specifies whether the script is enabled. The script runs only if enabled.

// Specifies whether the script is enabled. The script runs only if enabled.

// The global init script ID.

// The name of the script

// The position of a global init script, where 0 represents the first global
// init script to run, 1 is the second global init script to run, and so on. If
// you omit the position for a new global init script, it gets the last
// position. It runs after all current global init scripts. Setting any value
// greater than the position of the last script is equivalent to the last
// position. For example, suppose there are three existing scripts with
// positions 0, 1 and 2. Any position value of 3 or greater puts the script in
// the last position (3) If an explicit position value conflicts with an
// existing script, your request succeeds. The original script at that position
// and all later scripts have their position incremented by 1.

// The position of a script, where 0 represents the first script to run, 1 is
// the second script to run, and so on.

// The position of a script, where 0 represents the first script to run, 1 is
// the second script to run, and so on. To move the script so that it runs
// first, set its position to 0. To move the script to the end, set it to any
// value greater or equal to the position of the last script. For example,
// suppose there are three existing scripts with positions 0, 1 and 2. Any
// position value of 2 or greater puts the script in the last position (2). If
// an explicit position value conflicts with an existing script, your request
// succeeds. The original script at that position and all later scripts have
// their position incremented by 1.
