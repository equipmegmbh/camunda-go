package camunda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

var (
	url  = "http://127.0.0.1"
	path = "engine-rest"
)

var (
	client Client
)

// Configure configures the camunda client.
func Configure(endpoint, context string) {
	url = endpoint
	path = context
}

// CreateDeployment creates a Deployment.
func CreateDeployment(ctx context.Context, tenant, name, filename string, content io.Reader) (*Deployment, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := multipart.NewWriter(buffer)

	var fw io.Writer
	var uri string
	var n int64
	var err error

	if err = writer.WriteField("tenant-id", tenant); err != nil {
		return nil, err
	}

	if err = writer.WriteField("deployment-source", name); err != nil {
		return nil, err
	}

	if err = writer.WriteField("deploy-changed-only", "false"); err != nil {
		return nil, err
	}

	if err = writer.WriteField("enable-duplicate-filtering", "true"); err != nil {
		return nil, err
	}

	header := make(textproto.MIMEHeader)

	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "data", filename))

	if fw, err = writer.CreatePart(header); err != nil {
		return nil, err
	}

	if n, err = io.Copy(fw, content); n < 0 || err != nil {
		return nil, err
	}

	ct := writer.FormDataContentType()

	if err = writer.Close(); err != nil {
		return nil, err
	}

	result := new(Deployment)

	uri = fmt.Sprintf("%s/%s/deployment/create", url, path)
	err = client.send(ctx, uri, http.MethodPost, ct, buffer, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitions queries for process definitions that fulfill given parameters. Parameters may
// be the properties of process definitions, such as the name, key or version. The size of the result
// set can be retrieved by using the GetProcessDefinitionsCount method.
func GetProcessDefinitions(ctx context.Context, tenantId string) ([]*ProcessDefinition, error) {
	var uri string
	var err error

	result := make([]*ProcessDefinition, 0)

	uri = fmt.Sprintf("%s/%s/process-definition?tenantIdIn=%s", url, path, tenantId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionsCount ...
func GetProcessDefinitionsCount(ctx context.Context) {

}

// GetProcessDefinition retrieves a process definition according to the ProcessDefinition interface in the engine.
func GetProcessDefinition(ctx context.Context, id string) (*ProcessDefinition, error) {
	var uri string
	var err error

	result := new(ProcessDefinition)

	uri = fmt.Sprintf("%s/%s/process-definition/%s", url, path, id)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionByKey retrieves a process definition according to the ProcessDefinition interface
// in the engine. Returns the latest version of the ProcessDefinition which belongs to no tenant.
func GetProcessDefinitionByKey(ctx context.Context, key string) (*ProcessDefinition, error) {
	var uri string
	var err error

	result := new(ProcessDefinition)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s", url, path, key)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionByTenant retrieves a process definition according to the ProcessDefinition interface
// in the engine. Returns the latest version of the ProcessDefinition for tenant.
func GetProcessDefinitionByTenant(ctx context.Context, key, tenantId string) (*ProcessDefinition, error) {
	var uri string
	var err error

	result := new(ProcessDefinition)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s/tenant-id/%s", url, path, key, tenantId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionXML retrieves the BPMN 2.0 XML of a process definition.
func GetProcessDefinitionXML(ctx context.Context, id string) (*ProcessDefinitionSource, error) {
	var uri string
	var err error

	result := new(ProcessDefinitionSource)

	uri = fmt.Sprintf("%s/%s/process-definition/%s/xml", url, path, id)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionXMLByKey retrieves the BPMN 2.0 XML of a process definition. Returns the XML for the
// latest version of the process definition which belongs to no tenant.
func GetProcessDefinitionXMLByKey(ctx context.Context, key string) (*ProcessDefinitionSource, error) {
	var uri string
	var err error

	result := new(ProcessDefinitionSource)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s/xml", url, path, key)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessDefinitionXMLByTenant retrieves the BPMN 2.0 XML of a process definition. Returns the XML for the
// latest version of the process definition for tenant.
func GetProcessDefinitionXMLByTenant(ctx context.Context, key, tenantId string) (*ProcessDefinitionSource, error) {
	var uri string
	var err error

	result := new(ProcessDefinitionSource)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s/tenant-id/%s/xml", url, path, key, tenantId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// StartProcessDefinition instantiates a given process definition. Process variables and business
// key may be supplied in the request body.
func StartProcessDefinition(ctx context.Context, id string, trigger *ProcessInstanceTrigger) (*ProcessInstance, error) {
	var reader io.Reader

	var uri string
	var payload []byte
	var err error

	if trigger == nil {
		goto exec
	}

	payload, err = json.Marshal(trigger)

	if err != nil {
		return nil, err
	}

	reader = bytes.NewReader(payload)

exec:
	result := new(ProcessInstance)

	uri = fmt.Sprintf("%s/%s/process-definition/%s/start", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", reader, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// StartProcessDefinitionByKey instantiates a given process definition. Process variables and business
// key may be supplied in the request body. Starts the latest version of the process definition which belongs to no tenant.
func StartProcessDefinitionByKey(ctx context.Context, key string, trigger *ProcessInstanceTrigger) (*ProcessInstance, error) {
	var reader io.Reader

	var uri string
	var payload []byte
	var err error

	if trigger == nil {
		goto exec
	}

	payload, err = json.Marshal(trigger)

	if err != nil {
		return nil, err
	}

	reader = bytes.NewReader(payload)

exec:
	result := new(ProcessInstance)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s/start", url, path, key)
	err = client.send(ctx, uri, http.MethodPost, "application/json", reader, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// StartProcessDefinitionByTenant instantiates a given process definition. Process variables and business
// key may be supplied in the request body. Starts the latest version of the process definition for tenant
func StartProcessDefinitionByTenant(ctx context.Context, key, tenantId string, trigger *ProcessInstanceTrigger) (*ProcessInstance, error) {
	var reader io.Reader

	var uri string
	var payload []byte
	var err error

	if trigger == nil {
		goto exec
	}

	payload, err = json.Marshal(trigger)

	if err != nil {
		return nil, err
	}

	reader = bytes.NewReader(payload)

exec:
	result := new(ProcessInstance)

	uri = fmt.Sprintf("%s/%s/process-definition/key/%s/tenant-id/%s/start", url, path, key, tenantId)
	err = client.send(ctx, uri, http.MethodPost, "application/json", reader, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetProcessInstances queries for process instances that fulfill given parameters. Parameters may be
// static as well as dynamic runtime properties of process instances. The size of the result set can
// be retrieved by using the GetProcessInstancesCount method.
func GetProcessInstances(ctx context.Context, tenantId string) ([]*ProcessInstance, error) {
	var uri string
	var err error

	result := make([]*ProcessInstance, 0)

	uri = fmt.Sprintf("%s/%s/process-instance?tenantIdIn=%s", url, path, tenantId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetProcessInstancesCount ...
func GetProcessInstancesCount(ctx context.Context) {

}

// GetTasks queries for tasks that fulfill a given filter. The size of the result set can be retrieved
// by using the GetTasksCount method.
func GetTasks(ctx context.Context, processInstanceId string) ([]*Task, error) {
	var uri string
	var err error

	result := make([]*Task, 0)

	uri = fmt.Sprintf("%s/%s/task?processInstanceId=%s", url, path, processInstanceId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetTasksCount ...
func GetTasksCount(ctx context.Context) {

}

// GetTasksHistory queries for historic tasks that fulfill the given parameters. The size of the result
// set can be retrieved by using the GetTasksHistoryCount method.
func GetTasksHistory(ctx context.Context, processInstanceId string) ([]*TaskHistory, error) {
	var uri string
	var err error

	result := make([]*TaskHistory, 0)

	uri = fmt.Sprintf("%s/%s/history/task?processInstanceId=%s", url, path, processInstanceId)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetTasksHistoryCount ...
func GetTasksHistoryCount(ctx context.Context) {

}

// GetTask retrieves a task by id.
func GetTask(ctx context.Context, id string) (*Task, error) {
	var uri string
	var err error

	result := new(Task)

	uri = fmt.Sprintf("%s/%s/task/%s", url, path, id)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// ClaimTask claims a task for a specific user.
func ClaimTask(ctx context.Context, id, userId string) error {
	var uri string
	var err error

	payload, err := json.Marshal(&map[string]string{"userId": userId})

	if err != nil {
		return err
	}

	uri = fmt.Sprintf("%s/%s/task/%s/claim", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", bytes.NewReader(payload), nil)

	return err
}

// DelegateTask delegates a task to another user.
func DelegateTask(ctx context.Context, id, userId string) error {
	var uri string
	var err error

	payload, err := json.Marshal(&map[string]string{"userId": userId})

	if err != nil {
		return err
	}

	uri = fmt.Sprintf("%s/%s/task/%s/delegate", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", bytes.NewReader(payload), nil)

	return err
}

// UnclaimTask resets a task’s assignee. If successful, the task is not assigned to a user.
func UnclaimTask(ctx context.Context, id string) error {
	var uri string
	var err error

	uri = fmt.Sprintf("%s/%s/task/%s/unclaim", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", nil, nil)

	return err
}

// CompleteTask completes a task and updates process variables.
func CompleteTask(ctx context.Context, id string, variables map[string]*Variable) error {
	var reader io.Reader

	var uri string
	var payload []byte
	var err error

	if variables == nil {
		goto exec
	}

	payload, err = json.Marshal(&map[string]interface{}{"variables": variables})

	if err != nil {
		return err
	}

	reader = bytes.NewReader(payload)

exec:
	uri = fmt.Sprintf("%s/%s/task/%s/complete", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", reader, nil)

	return err
}

// ResolveTask resolves a task and updates execution variables.
// Resolving a task marks that the assignee is done with the task delegated to them, and that it
// can be sent back to the owner. Can only be executed when the task has been delegated. The assignee
// will be set to the owner, who performed the delegation.
func ResolveTask(ctx context.Context, id string, variables map[string]*Variable) error {
	var reader io.Reader

	var uri string
	var payload []byte
	var err error

	if variables == nil {
		goto exec
	}

	payload, err = json.Marshal(&map[string]interface{}{"variables": variables})

	if err != nil {
		return err
	}

	reader = bytes.NewReader(payload)

exec:
	uri = fmt.Sprintf("%s/%s/task/%s/resolve", url, path, id)
	err = client.send(ctx, uri, http.MethodPost, "application/json", reader, nil)

	return err
}

// GetTenants query for a list of tenants using a list of parameters. The size of the result
// set can be retrieved by using the GetTenantsCount method.
func GetTenants(ctx context.Context) ([]*Tenant, error) {
	var uri string
	var err error

	result := make([]*Tenant, 0)

	uri = fmt.Sprintf("%s/%s/tenant", url, path)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// GetTenantsCount ...
func GetTenantsCount(ctx context.Context) {

}

// GetTenant retrieves a Tenant.
func GetTenant(ctx context.Context, id string) (*Tenant, error) {
	var uri string
	var err error

	result := new(Tenant)

	uri = fmt.Sprintf("%s/%s/tenant/%s", url, path, id)
	err = client.send(ctx, uri, http.MethodGet, "application/json", nil, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// CreateTenant creates a new Tenant.
func CreateTenant(ctx context.Context, tenant *Tenant) error {
	var uri string
	var err error

	payload, err := json.Marshal(tenant)

	if err != nil {
		return err
	}

	uri = fmt.Sprintf("%s/%s/tenant/create", url, path)
	err = client.send(ctx, uri, http.MethodPost, "application/json", bytes.NewReader(payload), nil)

	return err
}
