// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

// all definitions in this file are in alphabetical order

type CreateScriptResponse struct {
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`
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
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	CreatedBy string `json:"created_by,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`

	Position int `json:"position,omitempty"`
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	UpdatedBy string `json:"updated_by,omitempty"`
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	CreatedBy string `json:"created_by,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`

	Position int `json:"position,omitempty"`

	Script string `json:"script,omitempty"`
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
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

// The name of the script

// The position of a global init script, where 0 represents the first script to
// run, 1 is the second script to run, in ascending order.
//
// If you omit the numeric position for a new global init script, it defaults to
// last position. It will run after all current scripts. Setting any value
// greater than the position of the last script is equivalent to the last
// position. Example: Take three existing scripts with positions 0, 1, and 2.
// Any position of (3) or greater puts the script in the last position. If an
// explicit position value conflicts with an existing script value, your request
// succeeds, but the original script at that position and all later scripts have
// their positions incremented by 1.

// The position of a script, where 0 represents the first script to run, 1 is
// the second script to run, in ascending order.

// The position of a script, where 0 represents the first script to run, 1 is
// the second script to run, in ascending order. To move the script to run
// first, set its position to 0.
//
// To move the script to the end, set its position to any value greater or equal
// to the position of the last script. Example, three existing scripts with
// positions 0, 1, and 2. Any position value of 2 or greater puts the script in
// the last position (2).
//
// If an explicit position value conflicts with an existing script, your request
// succeeds, but the original script at that position and all later scripts have
// their positions incremented by 1.
