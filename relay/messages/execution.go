package messages

import (
	"strings"
)

// ExecutionRequest is a request to execute a command
// as part of a Cog pipeline
type ExecutionRequest struct {
	Options       map[string]interface{} `json:"options"`
	Args          []interface{}          `json:"args"`
	CogEnv        map[string]interface{} `json:"cog_env"`
	Command       string                 `json:"command"`
	CommandConfig map[string]interface{} `json:"command_config"`
	ReplyTo       string                 `json:"reply_to"`
	Requestor     ChatUser               `json:"requestor"`
	User          CogUser                `json:"user"`
}

// ChatUser contains chat information about the submittor
type ChatUser struct {
	ID       string `json:"id"`
	Handle   string `json:"handle"`
	Provider string `json:"provider"`
}

// CogUser contains Cog user information about the submittor
type CogUser struct {
	ID        string `json:"id"`
	Email     string `json:"email_address"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// ExecutionResponse contains the results of executing a command
type ExecutionResponse struct {
	Room          string      `json:"room"`
	Bundle        string      `json:"bundle"`
	Status        string      `json:"status"`
	StatusMessage string      `json:"status_message"`
	Template      string      `json:"template"`
	Body          interface{} `json:"body"`
	IsJSON        bool
}

func (er *ExecutionRequest) BundleName() string {
	return strings.SplitN(er.Command, ":", 2)[0]
}

func (er *ExecutionRequest) CommandName() string {
	return strings.SplitN(er.Command, ":", 2)[1]
}

func (er *ExecutionRequest) PipelineID() string {
	return strings.SplitN(er.ReplyTo, "/", 5)[3]
}
