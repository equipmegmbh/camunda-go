package camunda

import (
	"fmt"
)

//goland:noinspection GoUnusedConst,GoNameStartsWithPackageName

type Error struct {
	// The type of exception.
	Type string `json:"type,omitempty"`

	// The error message.
	Message string `json:"message,omitempty"`
}

type Fetch struct {
	// Mandatory. The id of the worker on which behalf tasks are fetched. The returned tasks are
	// locked for that worker and can only be completed when providing the same worker id.
	WorkerId string `json:"workerId,omitempty"`

	// Mandatory. The maximum number of tasks to return.
	MaxTasks int `json:"maxTasks,omitempty"`

	// A boolean value, which indicates whether the task should be fetched based on its priority or arbitrarily.
	UsePriority bool `json:"usePriority,omitempty"`

	// The Long Polling timeout in milliseconds.
	// Note: The value cannot be set larger than 1.800.000 milliseconds (corresponds to 30 minutes).
	AsyncResponseTimeout int `json:"asyncResponseTimeout,omitempty"`

	// The topics for which external tasks should be fetched.
	Topics []*Topic `json:"topics,omitempty"`
}

type Task struct {
	// The id of the external task.
	Id string `json:"id,omitempty"`

	// The id of the activity that this external task belongs to.
	ActivityId string `json:"activityId,omitempty"`

	// The id of the activity instance that the external task belongs to.
	ActivityInstanceId string `json:"activityInstanceId,omitempty"`

	// The full error message submitted with the latest reported failure executing this task;
	// null if no failure was reported previously or if no error message was submitted.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The error details submitted with the latest reported failure executing this task.
	// null if no failure was reported previously or if no error details was submitted.
	ErrorDetails string `json:"errorDetails,omitempty"`

	// The id of the execution that the external task belongs to.
	ExecutionId string `json:"executionId,omitempty"`

	// The date that the task's most recent lock expires or has expired.
	LockExpirationTime string `json:"lockExpirationTime,omitempty"`

	// The id of the process definition the external task is defined in.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// The key of the process definition the external task is defined in.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// The version tag of the process definition the external task is defined in.
	ProcessDefinitionVersionTag string `json:"processDefinitionVersionTag,omitempty"`

	// The id of the process instance the external task belongs to.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`

	// The id of the tenant the external task belongs to.
	TenantId string `json:"tenantId,omitempty"`

	// The number of retries the task currently has left.
	Retries int `json:"retries,omitempty"`

	// Whether the process instance the external task belongs to is suspended.
	Suspended bool `json:"suspended,omitempty"`

	// The id of the worker that possesses or possessed the most recent lock.
	WorkerId string `json:"workerId,omitempty"`

	// The priority of the external task.
	Priority int `json:"priority,omitempty"`

	// The topic name of the external task.
	TopicName string `json:"topicName,omitempty"`

	// The business key of the process instance the external task belongs to.
	BusinessKey string `json:"businessKey,omitempty"`

	// The properties for each of the requested variables.
	Variables map[string]*Variable `json:"variables,omitempty"`
}

type Topic struct {
	// The topic's name. (Mandatory)
	TopicName string `json:"topicName,omitempty"`

	// The duration to lock the external tasks for in milliseconds. (Mandatory)
	LockDuration int `json:"lockDuration,omitempty"`

	// A JSON array of String values that represent variable names. For each result task belonging to
	// this topic, the given variables are returned as well if they are accessible from the external task's
	// execution. If not provided - all variables will be fetched.
	Variables []string `json:"variables,omitempty"`

	// If true only local variables will be fetched.
	LocalVariables bool `json:"localVariables,omitempty"`

	// A String value which enables the filtering of tasks based on process instance business key.
	BusinessKey string `json:"businessKey,omitempty"`

	// Filter tasks based on process definition id.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// Filter tasks based on process definition ids.
	ProcessDefinitionIdIn []string `json:"processDefinitionIdIn,omitempty"`

	// Filter tasks based on process definition key.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// Filter tasks based on process definition keys.
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`

	// Filter tasks based on process definition version tag.
	ProcessDefinitionVersionTag string `json:"processDefinitionVersionTag,omitempty"`

	// Filter tasks without tenant id.
	WithoutTenantId string `json:"withoutTenantId,omitempty"`

	// Filter tasks based on tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
}

type Variable struct {
	// The value type of the variable.
	Type string `json:"type,omitempty"`

	// The variable's value.
	Value interface{} `json:"value,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("type %s, message: %s", e.Type, e.Message)
}
