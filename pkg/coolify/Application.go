package coolify

type Application struct {
	Id                        string `json:"id"`
	UUID                      string `json:"uuid"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	PrivateKeyId              string `json:"private_key_id"`
	PrivateKeyUUID            string `json:"private_key_uuid"`
	ProjectId                 string `json:"project_id"`
	ProjectUUID               string `json:"project_uuid"`
	ServerId                  string `json:"server_id"`
	ServerUUID                string `json:"server_uuid"`
	EnvironmentName           string `json:"environment_name"`
	GitRepository             string `json:"git_repository"`
	GitBranch                 string `json:"git_branch"`
	BuildPack                 string `json:"build_pack"`
	PortsExposes              string `json:"ports_exposes"`
	PortsMappings             string `json:"ports_mappings"`
	BaseDirectory             string `json:"base_directory"`
	LimitsMemory              string `json:"limits_memory"`
	LimitsCpus                string `json:"limits_cpus"`
	CustomDockerRunOptions    string `json:"custom_docker_run_options"`
	PostDeploymentCommand     string `json:"post_deployment_command"`
	PreDeploymentCommand      string `json:"pre_deployment_command"`
	ManualWebhookSecretGithub string `json:"manual_webhook_secret_github"`
	ManualWebhookSecretGitlab string `json:"manual_webhook_secret_gitlab"`
}

func (app *Application) GetUUID() string {
	return app.UUID
}

func (app *Application) SetUUID(uuid string) {
	app.UUID = uuid
}

func (app *Application) GetIdentifier() string {
	return app.Id
}

func (app *Application) ToSaveState() (map[string]interface{}, error) {
	m, err := structToMap(app)
	if err != nil {
		return nil, err
	}

	hashValues(m, []string{
		"manual_webhook_secret_github",
		"manual_webhook_secret_gitlab",
	})

	return m, nil
}

func (app *Application) ToCreatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "server_id", "project_id", "private_key_id"}
	return structToPayload(app, ignore_attrs)
}

func (app *Application) ToUpdatePayload() ([]byte, error) {
	ignore_attrs := []string{
		"id",
		"uuid",
		"project_id",
		"server_id",
		"private_key_id",
		"project_uuid",
		"server_uuid",
		"private_key_uuid",
		"environment_name",
	}
	return structToPayload(app, ignore_attrs)
}

func (app *Application) BuildNewFromCurrentState(state *State) Resource {
	copy := *app
	copy.ServerUUID = state.serverUUID(app.ServerId)
	copy.ProjectUUID = state.projectUUID(app.ProjectId)
	copy.PrivateKeyUUID = state.privateKeyUUID(app.PrivateKeyId)
	return &copy
}
