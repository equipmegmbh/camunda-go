package camunda

import (
	"fmt"
)

//goland:noinspection GoUnusedConst,GoNameStartsWithPackageName

type CaseDefinition struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// The key of the process definition.
	Key string `json:"key,omitempty"`

	// The category of the process definition.
	Category string `json:"category,omitempty"`

	// The name of the process definition.
	Name string `json:"name,omitempty"`

	// The version of the process definition that the engine assigned to it.
	Version int `json:"version,omitempty"`

	// The file name of the process definition.
	Resource string `json:"resource,omitempty"`

	// The deployment id of the process definition.
	DeploymentId string `json:"deploymentId,omitempty"`

	// The tenant id of the process definition.
	TenantId string `json:"tenantId,omitempty"`

	// History time to live value of the process definition.
	HistoryTimeToLive int `json:"historyTimeToLive,omitempty"`
}

type Deployment struct {
	// The id of the deployment.
	Id string `json:"id,omitempty"`

	// The name of the deployment.
	Name string `json:"name,omitempty"`

	// The source of the deployment.
	Source string `json:"source,omitempty"`

	// The tenant id of the deployment.
	TenantId string `json:"tenantId,omitempty"`

	// The time when the deployment was created.
	DeploymentTime string `json:"deploymentTime,omitempty"`

	// An object containing a property for each of the
	// process definitions, which are successfully deployed with that deployment.
	DeployedProcessDefinitions map[string]*ProcessDefinition `json:"deployedProcessDefinitions,omitempty"`

	// An object containing a property for each of the case
	// definitions, which are successfully deployed with that deployment.
	DeployedCaseDefinitions map[string]*CaseDefinition `json:"deployedCaseDefinitions,omitempty"`

	// An object containing a property for each of the decision
	// definitions, which are successfully deployed with that deployment.
	DeployedDecisionDefinitions map[string]*DecisionDefinition `json:"deployedDecisionDefinitions,omitempty"`

	// An object containing a property for each of the decision
	// requirements definitions, which are successfully deployed with that deployment.
	DeployedDecisionRequirementsDefinitions map[string]*DecisionRequirementsDefinition `json:"deployedDecisionRequirementsDefinitions,omitempty"`

	// Link to the newly created deployment with method, href and rel.
	Links []*Link `json:"links,omitempty"`
}

type DecisionDefinition struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// The key of the process definition.
	Key string `json:"key,omitempty"`

	// The category of the process definition.
	Category string `json:"category,omitempty"`

	// The name of the process definition.
	Name string `json:"name,omitempty"`

	// The version of the process definition that the engine assigned to it.
	Version int `json:"version,omitempty"`

	// The file name of the process definition.
	Resource string `json:"resource,omitempty"`

	// The deployment id of the process definition.
	DeploymentId string `json:"deploymentId,omitempty"`

	// The id of the decision requirements definition this decision definition belongs to.
	DecisionRequirementsDefinitionId string `json:"decisionRequirementsDefinitionId,omitempty"`

	// The key of the decision requirements definition this decision definition belongs to.
	DecisionRequirementsDefinitionKey string `json:"decisionRequirementsDefinitionKey,omitempty"`

	// The tenant id of the process definition.
	TenantId string `json:"tenantId,omitempty"`

	// The version tag of the process definition.
	VersionTag string `json:"versionTag,omitempty"`

	// History time to live value of the process definition.
	HistoryTimeToLive int `json:"historyTimeToLive,omitempty"`
}

type DecisionRequirementsDefinition struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// The key of the process definition.
	Key string `json:"key,omitempty"`

	// The category of the process definition.
	Category string `json:"category,omitempty"`

	// The name of the process definition.
	Name string `json:"name,omitempty"`

	// The version of the process definition that the engine assigned to it.
	Version int `json:"version,omitempty"`

	// The file name of the process definition.
	Resource string `json:"resource,omitempty"`

	// The deployment id of the process definition.
	DeploymentId string `json:"deploymentId,omitempty"`

	// The tenant id of the process definition.
	TenantId string `json:"tenantId,omitempty"`
}

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

type Link struct {
	// The link method.
	Method string `json:"method,omitempty"`

	// The link href.
	Ref string `json:"href,omitempty"`

	// The link relation.
	Rel string `json:"rel,omitempty"`
}

type ProcessDefinition struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// The key of the process definition.
	Key string `json:"key,omitempty"`

	// The category of the process definition.
	Category string `json:"category,omitempty"`

	// The description of the process definition.
	Description string `json:"description,omitempty"`

	// The name of the process definition.
	Name string `json:"name,omitempty"`

	// The version of the process definition that the engine assigned to it.
	Version int `json:"version,omitempty"`

	// The file name of the process definition.
	Resource string `json:"resource,omitempty"`

	// The deployment id of the process definition.
	DeploymentId string `json:"deploymentId,omitempty"`

	// The file name of the process definition diagram, if it exists.
	Diagram string `json:"diagram,omitempty"`

	// A flag indicating whether the definition is suspended or not.
	Suspended bool `json:"suspended,omitempty"`

	// The tenant id of the process definition.
	TenantId string `json:"tenantId,omitempty"`

	// The version tag of the process definition.
	VersionTag string `json:"versionTag,omitempty"`

	// History time to live value of the process definition.
	HistoryTimeToLive int `json:"historyTimeToLive,omitempty"`

	// A flag indicating whether the process definition is startable in Tasklist or not.
	StartableInTasklist bool `json:"startableInTasklist,omitempty"`
}

type ProcessDefinitionSource struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// An escaped XML string containing the XML that this definition was deployed with.
	Content string `json:"bpmn20Xml,omitempty"`
}

type ProcessInstance struct {
	// The id of the process definition.
	Id string `json:"id,omitempty"`

	// The id of the process definition.
	DefinitionId string `json:"definitionId,omitempty"`

	// The business key of the process instance.
	BusinessKey string `json:"businessKey,omitempty"`

	// The case instance id of the process instance.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// The tenant id of the process instance.
	TenantId string `json:"tenantId,omitempty"`

	// A flag indicating whether the instance is still running or not.
	Ended bool `json:"ended,omitempty"`

	// A flag indicating whether the instance is suspended or not.
	Suspended bool `json:"suspended,omitempty"`

	// An array containing links to interact with the instance.
	Links []*Link `json:"links,omitempty"`

	// An object containing a property for each of the latest variables.
	Variables map[string]*Variable `json:"variables,omitempty"`
}

type ProcessInstanceTrigger struct {
	// The business key the process instance is to be initialized with. The business key uniquely
	// identifies the process instance in the context of the given process definition.
	BusinessKey string `json:"businessKey,omitempty"`

	// The case instance id the process instance is to be initialized with.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// Skip execution listener invocation for activities that are started
	// or ended as part of this request.
	SkipCustomListeners bool `json:"skipCustomListeners,omitempty"`

	// Skip execution of input/output variable mappings for activities that are
	// started or ended as part of this request.
	SkipIoMappings bool `json:"skipIoMappings,omitempty"`

	// Indicates if the variables, which was used by the process instance during
	// execution, should be returned. Default value: false
	WithVariablesInReturn bool `json:"withVariablesInReturn,omitempty"`

	// An array of instructions that specify which activities to start the process
	// instance at. If this property is omitted, the process instance starts at its default blank start event.
	StartInstructions []*StartInstruction `json:"startInstructions,omitempty"`

	// An object containing a property for each of the latest variables.
	Variables map[string]*Variable `json:"variables,omitempty"`
}

type StartInstruction struct {
	// Mandatory. One of the following values: 'startBeforeActivity', 'startAfterActivity', 'startTransition'.
	// A 'startBeforeActivity' instruction requests to start execution before entering a given
	// activity. A 'startAfterActivity' instruction requests to start at the single outgoing sequence
	// flow of a given activity. A 'startTransition' instruction requests to execute a specific sequence flow.
	Type string `json:"type,omitempty"`

	// Specifies the activity the instruction targets. Can be used with instructions
	// of types 'startBeforeActivity' and 'startAfterActivity'.
	ActivityId string `json:"activityId,omitempty"`

	// Specifies the sequence flow to start. Can be used with
	// instructions of types 'startTransition'.
	TransitionId string `json:"transitionId,omitempty"`

	// An object containing variable key-value pairs. Can be used with instructions
	// of type 'startBeforeActivity', 'startAfterActivity', and 'startTransition'.
	Variables map[string]*Variable `json:"variables,omitempty"`
}

type Task struct {
	// The task id.
	Id string `json:"id,omitempty"`

	// The task name.
	Name string `json:"name,omitempty"`

	// The assignee's id.
	Assignee string `json:"assignee,omitempty"`

	// The owner's id.
	Owner string `json:"owner,omitempty"`

	// The date the task was created on. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	Created string `json:"created,omitempty"`

	// The task's due date. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	Due string `json:"due,omitempty"`

	// The follow-up date for the task. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	FollowUp string `json:"followUp,omitempty"`

	// The task's delegation state. Possible values are PENDING and RESOLVED.
	DelegationState string `json:"delegationState,omitempty"`

	// The task's description.
	Description string `json:"description,omitempty"`

	// The id of the execution the task belongs to.
	ExecutionId string `json:"executionId,omitempty"`

	// The id the parent task, if this task is a subtask.
	ParentTaskId string `json:"parentTaskId,omitempty"`

	// The task's priority.
	Priority int `json:"priority,omitempty"`

	// The id of the process definition the task belongs to.
	ProcessDefinitionId int `json:"processDefinitionId,omitempty"`

	// The id of the process instance the task belongs to.
	ProcessInstanceId bool `json:"processInstanceId,omitempty"`

	// The id of the case execution the task belongs to.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`

	// The id of the case definition the task belongs to.
	CaseDefinitionId int `json:"caseDefinitionId,omitempty"`

	// The id of the case instance the task belongs to.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// The task's key.
	TaskDefinitionKey string `json:"taskDefinitionKey,omitempty"`

	// Whether the task belongs to a process instance that is suspended.
	Suspended bool `json:"suspended,omitempty"`

	// If not null, the form key for the task.
	FormKey string `json:"formKey,omitempty"`

	// If not null, the tenant id of the task.
	TenantId string `json:"tenantId,omitempty"`
}

type TaskHistory struct {
	// The task id.
	Id string `json:"id,omitempty"`

	// The key of the process definition the task belongs to.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// The id of the process definition the task belongs to.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// The id of the process instance the task belongs to.
	ProcessInstanceId string `json:"processInstanceId	,omitempty"`

	// The id of the execution the task belongs to.
	ExecutionId string `json:"executionId,omitempty"`

	// The key of the case definition the task belongs to.
	CaseDefinitionKey string `json:"caseDefinitionKey,omitempty"`

	// The id of the case definition the task belongs to.
	CaseDefinitionId string `json:"caseDefinitionId,omitempty"`

	// The id of the case instance the task belongs to.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// The id of the case execution the task belongs to.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`

	// The id of the activity that this object is an instance of.
	ActivityInstanceId string `json:"activityInstanceId,omitempty"`

	// The task name.
	Name string `json:"name,omitempty"`

	// The task's description.
	Description string `json:"description,omitempty"`

	// The task's delete reason.
	DeleteReason string `json:"deleteReason,omitempty"`

	// The owner's id.
	Owner string `json:"owner,omitempty"`

	// The assignee's id.
	Assignee string `json:"assignee,omitempty"`

	// The time the task was started. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	StartTime string `json:"startTime,omitempty"`

	// The time the task ended. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	EndTime string `json:"endTime,omitempty"`

	// The time the task took to finish (in milliseconds).
	Duration int `json:"duration,omitempty"`

	// The task's key.
	TaskDefinitionKey string `json:"taskDefinitionKey,omitempty"`

	// The task's priority.
	Priority int `json:"priority,omitempty"`

	// The task's due date. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	Due string `json:"due,omitempty"`

	// The id of the parent task, if this task is a subtask.
	ParentTaskId string `json:"parentTaskId,omitempty"`

	// The follow-up date for the task. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	FollowUp string `json:"followUp,omitempty"`

	// The tenant id of the task instance.
	TenantId string `json:"tenantId,omitempty"`

	// The time after which the task should be removed by the History Cleanup job.
	// Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	RemovalTime string `json:"removalTime,omitempty"`

	// The process instance id of the root process instance that initiated the
	// process containing this task.
	RootProcessInstanceId string `json:"rootProcessInstanceId,omitempty"`
}

type Tenant struct {
	// The id of the tenant.
	Id string `json:"id,omitempty"`

	// The name of the tenant.
	Name string `json:"name,omitempty"`
}

type Token struct {
	// ...
	Kind string `json:"token_type,omitempty"`

	// ...
	Scope string `json:"scope,omitempty"`

	// ...
	ExpiresIn int32 `json:"expires_in,string"`

	// ...
	ExpiresOn int64 `json:"expires_on,string"`

	// ...
	AccessToken string `json:"access_token,omitempty"`

	// ...
	RefreshToken string `json:"refresh_token,omitempty"`
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
	// Indicates whether the variable should be a local variable or not.
	Local bool `json:"local,omitempty"`

	// The value type of the variable.
	Type string `json:"type,omitempty"`

	// The variable's value.
	Value interface{} `json:"value,omitempty"`

	// An object containing additional, value-type-dependent properties.
	ValueInfo interface{} `json:"valueInfo,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("type %s, message: %s", e.Type, e.Message)
}
