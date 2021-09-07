package camunda

import (
	"fmt"
)

//goland:noinspection GoUnusedConst,GoNameStartsWithPackageName

const (
	TimeLayout = "2006-01-02T15:04:05.000-07:00"
)

type Batch struct {
	// The id of the batch.
	Id string `json:"id,omitempty"`

	// The type of the batch.
	Type string `json:"type,omitempty"`

	// The total jobs of a batch is the number of batch execution jobs required to complete the batch.
	TotalJobs int `json:"totalJobs,omitempty"`

	// The number of batch execution jobs already created by the seed job.
	JobsCreated int `json:"jobsCreated,omitempty"`

	// The number of batch execution jobs created per seed job invocation. The
	// batch seed job is invoked until it has created all batch execution jobs
	// required by the batch (see totalJobs property).
	BatchJobsPerSeed int `json:"batchJobsPerSeed,omitempty"`

	// Every batch execution job invokes the command executed by the batch invocationsPerBatchJob
	// times. E.g., for a process instance migration batch this specifies the number of process
	// instances which are migrated per batch execution job.
	InvocationsPerBatchJob int `json:"invocationsPerBatchJob,omitempty"`

	// The job definition id for the seed jobs of this batch.
	SeedJobDefinitionId string `json:"seedJobDefinitionId,omitempty"`

	// The job definition id for the monitor jobs of this batch.
	MonitorJobDefinitionId string `json:"monitorJobDefinitionId,omitempty"`

	// The job definition id for the batch execution jobs of this batch.
	BatchJobDefinitionId string `json:"batchJobDefinitionId,omitempty"`

	// Indicates whether this batch is suspended or not.
	Suspended bool `json:"suspended,omitempty"`

	// The tenant id of the batch.
	TenantId string `json:"tenantId,omitempty"`

	// The id of the user that created the batch.
	CreateUserId string `json:"createUserId,omitempty"`
}

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

type Comment struct {
	// The id of the task comment.
	Id string `json:"id,omitempty"`

	// The id of the user who created the comment.
	UserId string `json:"userId,omitempty"`

	// The id of the task to which the comment belongs.
	TaskId string `json:"taskId,omitempty"`

	// The id of the process instance the comment is related to.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`

	// The time when the comment was created.
	Time string `json:"time,omitempty"`

	// The content of the comment.
	Message string `json:"message,omitempty"`

	// Link to the newly created task comment with method, href and rel.
	Links []*Link `json:"links,omitempty"`

	// The time after which the comment should be removed by the History Cleanup
	// job. Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	RemovalTime string `json:"removalTime,omitempty"`

	// The process instance id of the root process instance that initiated the process containing the task.
	RootProcessInstanceId string `json:"rootProcessInstanceId,omitempty"`
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

type Instruction struct {
	// Mandatory. One of the following values: startBeforeActivity, startAfterActivity, startTransition.
	// A startBeforeActivity instruction requests to enter a given activity. A startAfterActivity instruction
	// requests to execute the single outgoing sequence flow of a given activity. A startTransition
	// instruction requests to execute a specific sequence flow.
	Type string `json:"type,omitempty"`

	// Can be used with instructions of types startBeforeActivity and startAfterActivity.
	// Specifies the activity the instruction targets.
	ActivityId string `json:"activityId,omitempty"`

	// Can be used with instructions of types startTransition. Specifies the sequence flow to start.
	TransitionId string `json:"transitionId,omitempty"`
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

type ProcessDefinitionStart struct {
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

type ProcessDefinitionRestart struct {
	// A list of process instance ids to restart.
	ProcessInstanceIds string `json:"processInstanceIds,omitempty"`

	// A historic process instance query. See ProcessInstanceQuery.
	HistoricProcessInstanceQuery *ProcessInstanceQuery `json:"historicProcessInstanceQuery,omitempty"`

	// Skip execution listener invocation for activities that are started as part of this request.
	SkipCustomListeners bool `json:"skipCustomListeners,omitempty"`

	// Skip execution of input/output variable mappings for activities that are started as part of this request.
	SkipIoMappings bool `json:"skipIoMappings,omitempty"`

	// Set the initial set of variables during restart. By default, the last set of variables is used.
	InitialVariables bool `json:"initialVariables,omitempty"`

	// Do not take over the business key of the historic process instance.
	WithoutBusinessKey bool `json:"withoutBusinessKey,omitempty"`

	// An array of instructions. The instructions are executed in the order they are in.
	Instructions []*Instruction `json:"instructions,omitempty"`
}

type ProcessInstance struct {
	// The id of the process instance.
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

type HistoricProcessInstance struct {
	// The id of the process instance.
	Id string `json:"id,omitempty"`

	// The process instance id of the root process instance that initiated the process.
	RootProcessInstanceId string `json:"rootProcessInstanceId,omitempty"`

	// The id of the parent process instance, if it exists.
	SuperProcessInstanceId string `json:"superProcessInstanceId,omitempty"`

	// The id of the parent case instance, if it exists.
	SuperCaseInstanceId string `json:"superCaseInstanceId,omitempty"`

	// The id of the parent case instance, if it exists.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// The name of the process definition that this process instance belongs to.
	ProcessDefinitionName string `json:"processDefinitionName,omitempty"`

	// The key of the process definition that this process instance belongs to.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// The version of the process definition that this process instance belongs to.
	ProcessDefinitionVersion int `json:"processDefinitionVersion,omitempty"`

	// The id of the process definition that this process instance belongs to.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// The business key of the process instance.
	BusinessKey string `json:"businessKey,omitempty"`

	// The time the instance was started. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	StartTime string `json:"startTime,omitempty"`

	// The time the instance ended. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	EndTime string `json:"endTime,omitempty"`

	// The time after which the instance should be removed by the History Cleanup job. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	RemovalTime string `json:"removalTime,omitempty"`

	// The time the instance took to finish (in milliseconds).
	DurationInMillis int `json:"durationInMillis,omitempty"`

	// The id of the user who started the process instance.
	StartUserId string `json:"startUserId,omitempty"`

	// The id of the initial activity that was executed (e.g., a start event).
	StartActivityId string `json:"startActivityId,omitempty"`

	// The provided delete reason in case the process instance was canceled during execution.
	DeleteReason string `json:"deleteReason,omitempty"`

	// The tenant id of the process instance.
	TenantId string `json:"tenantId,omitempty"`

	//last state of the process instance, possible values are:
	//ACTIVE - running process instance
	//SUSPENDED - suspended process instances
	//COMPLETED - completed through normal end event
	//EXTERNALLY_TERMINATED - terminated externally, for instance through REST API
	//INTERNALLY_TERMINATED - terminated internally, for instance by terminating boundary event
	State string `json:"state,omitempty"`
}

type ProcessInstanceQuery struct {
	// The id of the process definition.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`

	// Filter by process instance ids. Must be an array process instance ids.
	ProcessInstanceIds []string `json:"processInstanceIds,omitempty"`

	// Filter by process instance business key.
	ProcessInstanceBusinessKey string `json:"processInstanceBusinessKey	,omitempty"`

	// Filter by process instance business key that the parameter is a substring of.
	ProcessInstanceBusinessKeyLike string `json:"processInstanceBusinessKeyLike,omitempty"`

	// Restrict the query to all process instances that are top level process instances.
	RootProcessInstances string `json:"rootProcessInstances,omitempty"`

	// Restrict query to all process instances that are sub process instances of
	// the given process instance. Takes a process instance id.
	SuperProcessInstanceId string `json:"superProcessInstanceId	,omitempty"`

	// Restrict query to one process instance that has a sub process instance with the given id.
	SubProcessInstanceId string `json:"subProcessInstanceId,omitempty"`

	// Restrict query to all process instances that are sub process instances of the given
	// case instance. Takes a case instance id.
	SuperCaseInstanceId string `json:"superCaseInstanceId,omitempty"`

	// Restrict query to one process instance that has a sub case instance with the given id.
	SubCaseInstanceId string `json:"subCaseInstanceId,omitempty"`

	// Restrict query to all process instances that are sub process instances of the
	// given case instance. Takes a case instance id.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// Filter by the process definition the instances run on.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// Filter by the key of the process definition the instances run on.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// Filter by a list of process definition keys. A process instance must have one of the given
	// process definition keys. Must be an array of strings.
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`

	// Exclude instances that belong to a set of process definitions. Must be an array
	// of process definition keys.
	ProcessDefinitionKeyNotIn []string `json:"processDefinitionKeyNotIn,omitempty"`

	// Filter by the name of the process definition the instances run on.
	ProcessDefinitionName string `json:"processDefinitionName,omitempty"`

	// Filter by process definition names that the parameter is a substring of.
	ProcessDefinitionNameLike string `json:"processDefinitionNameLike,omitempty"`

	// Only include finished process instances. This flag includes all process instances
	// that are completed or terminated. Value may only be true, as false is the default behavior.
	Finished bool `json:"finished,omitempty"`

	// Only include unfinished process instances. Value may only be true, as false is the default behavior.
	Unfinished bool `json:"unfinished,omitempty"`

	// Only include process instances which have an incident. Value may only be
	// true, as false is the default behavior.
	WithIncidents bool `json:"withIncidents,omitempty"`

	// Only include process instances which have a root incident. Value may only be
	// true, as false is the default behavior.
	WithRootIncidents bool `json:"withRootIncidents,omitempty"`

	// Filter by the incident type. Valid values are failedJob or failedExternalTask.
	IncidentType string `json:"incidentType,omitempty"`

	// Only include process instances which have an incident in status either open or resolved.
	// To get all process instances, use the query parameter withIncidents.
	IncidentStatus string `json:"incidentStatus,omitempty"`

	// Filter by the incident message. Exact match.
	IncidentMessage string `json:"incidentMessage,omitempty"`

	// Filter by the incident message that the parameter is a substring of.
	IncidentMessageLike string `json:"incidentMessageLike,omitempty"`

	// Only include process instances that were started by the given user.
	StartedBy string `json:"startedBy,omitempty"`

	// Restrict to instances that were started before the given date. By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	StartedBefore string `json:"startedBefore,omitempty"`

	// Restrict to instances that were started after the given date. By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	StartedAfter string `json:"startedAfter,omitempty"`

	// Restrict to instances that were finished before the given date. By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	FinishedBefore string `json:"finishedBefore,omitempty"`

	// Restrict to instances that were finished after the given date. By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	FinishedAfter string `json:"finishedAfter,omitempty"`

	// Filter by a list of tenant ids. A process instance must have one of the given
	// tenant ids. Must be an array of strings.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`

	// Only include historic process instances which belong to no tenant.
	// Value may only be true, as false is the default behavior.
	WithoutTenantId bool `json:"withoutTenantId,omitempty"`

	// An array of QueryVariable to only include process instances that
	// have/had variables with certain values.
	Variables []*QueryVariable `json:"variables,omitempty"`

	// Match all variable names provided in variables case-insensitively. If
	// set to true variableName and variablename are treated as equal.
	VariableNamesIgnoreCase bool `json:"variableNamesIgnoreCase,omitempty"`

	// Match all variable values provided in variables case-insensitively. If
	// set to true variableValue and variablevalue are treated as equal.
	VariableValuesIgnoreCase bool `json:"variableValuesIgnoreCase,omitempty"`

	// An array of criteria to sort the result by (see Sort). Each element of the array is a object that
	// specifies one ordering. The position in the array identifies the rank of an ordering,
	// i.e., whether it is primary, secondary, etc.
	Sorting []*Sort `json:"sorting,omitempty"`

	// Restrict to instances that executed an activity before the given date (inclusive). By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedActivityBefore string `json:"executedActivityBefore,omitempty"`

	// Restrict to instances that executed an activity after the given date (inclusive). By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedActivityAfter string `json:"executedActivityAfter,omitempty"`

	// Restrict to instances that executed an activity with one of given ids.
	ExecutedActivityIdIn string `json:"executedActivityIdIn,omitempty"`

	// Restrict to instances that have an active activity with one of given ids.
	ActiveActivityIdIn string `json:"activeActivityIdIn,omitempty"`

	// Restrict to instances that executed a job before the given date (inclusive). By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedJobBefore string `json:"executedJobBefore,omitempty"`

	// Restrict to instances that executed a job after the given date (inclusive). By default,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedJobAfter string `json:"executedJobAfter,omitempty"`

	// Restrict to instances that are active.
	Active bool `json:"active,omitempty"`

	// Restrict to instances that are suspended.
	Suspended bool `json:"suspended,omitempty"`

	// Restrict to instances that are completed.
	Completed bool `json:"completed,omitempty"`

	// Restrict to instances that are externally terminated.
	ExternallyTerminated bool `json:"externallyTerminated,omitempty"`

	// Restrict to instances that are internally terminated.
	InternallyTerminated bool `json:"internallyTerminated,omitempty"`

	// An array of nested process instance queries with OR semantics. A process instance
	// matches a nested query if it fulfills at least one of the query's predicates. With multiple
	// nested queries, a process instance must fulfill at least one predicate of each query (Conjunctive Normal Form).
	//
	// All process instance query properties can be used except for: sorting.
	OrQueries []*ProcessInstanceQuery `json:"orQueries,omitempty"`
}

type QueryVariable struct {
	// The variable name.
	Name string `json:"name,omitempty"`

	// The comparison operator to be used. Valid operator values are:
	// eq - equal to; neq - not equal to; gt - greater than;
	// gteq - greater than or equal to; lt - lower than;
	// lteq - lower than or equal to; like.
	Operator string `json:"operator,omitempty"`

	// The variable value. May be String, Number or Boolean
	Value interface{} `json:"value,omitempty"`
}

type Sort struct {
	// Mandatory. Sorts the results lexicographically by a given
	// criterion. Valid values are instanceId, definitionId, definitionKey,
	// definitionName, definitionVersion, businessKey, startTime, endTime, duration and tenantId.
	SortBy string `json:"sortBy,omitempty"`

	// Mandatory. Sorts the results in a given order. Values may
	// be asc for ascending order or desc for descending order.
	SortOrder string `json:"sortOrder,omitempty"`
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
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// The id of the process instance the task belongs to.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`

	// The id of the case execution the task belongs to.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`

	// The id of the case definition the task belongs to.
	CaseDefinitionId string `json:"caseDefinitionId,omitempty"`

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

type UserOperationLog struct {
	// The unique identifier of this log entry.
	Id string `json:"id,omitempty"`

	// The user who performed this operation.
	UserId string `json:"userId,omitempty"`

	// Timestamp of this operation.
	Timestamp string `json:"timestamp,omitempty"`

	// The unique identifier of this operation. A composite operation that changes
	// multiple properties has a common operationId.
	OperationId string `json:"operationId,omitempty"`

	// The type of this operation, e.g., Assign, Claim and so on.
	OperationType string `json:"operationType,omitempty"`

	// The type of the entity on which this operation was executed, e.g., Task or Attachment.
	EntityType string `json:"entityType,omitempty"`

	// The name of the category this operation was associated with, e.g., TaskWorker or Admin.
	Category string `json:"category,omitempty"`

	// An arbitrary annotation set by a user for auditing reasons.
	Annotation string `json:"annotation,omitempty"`

	// The property changed by this operation.
	Property string `json:"property,omitempty"`

	// The original value of the changed property.
	OrgValue string `json:"orgValue,omitempty"`

	// The new value of the changed property.
	NewValue string `json:"newValue,omitempty"`

	// If not null, the operation is restricted to entities in relation to this deployment.
	DeploymentId string `json:"deploymentId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this process definition.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`

	// If not null, the operation is restricted to entities in relation to process definitions with this key.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`

	// If not null, the operation is restricted to entities in relation to this process instance.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this execution.
	ExecutionId string `json:"executionId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this case definition.
	CaseDefinitionId string `json:"caseDefinitionId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this case instance.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this case execution.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this task.
	TaskId string `json:"taskId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this job.
	JobId string `json:"jobId,omitempty"`

	// If not null, the operation is restricted to entities in relation to this job definition.
	JobDefinitionId string `json:"jobDefinitionId,omitempty"`

	// The time after which the entry should be removed by the History Cleanup job.
	// Default format* yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	RemovalTime string `json:"removalTime,omitempty"`

	// The process instance id of the root process instance that initiated the process containing this entry.
	RootProcessInstanceId string `json:"rootProcessInstanceId,omitempty"`
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
