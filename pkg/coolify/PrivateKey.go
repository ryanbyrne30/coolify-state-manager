package coolify

type PrivateKey struct {
	Id          string `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PrivateKey  string `json:"private_key"`
}

func (k *PrivateKey) GetUUID() string {
	return k.UUID
}

func (k *PrivateKey) SetUUID(uuid string) {
	k.UUID = uuid
}

func (k *PrivateKey) GetIdentifier() string {
	return k.Id
}

func (k *PrivateKey) ToSaveState() (map[string]interface{}, error) {
	m, err := structToMap(k)
	if err != nil {
		return nil, err
	}

	hashValues(m, []string{
		"private_key",
	})

	return m, nil
}

func (k *PrivateKey) ToCreatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "key_hash"}
	return structToPayload(k, ignore_attrs)
}

func (k *PrivateKey) ToUpdatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "key_hash"}
	return structToPayload(k, ignore_attrs)
}

func (k *PrivateKey) BuildNewFromCurrentState(state *State) Resource {
	copy := *k
	return &copy
}

func GetPrivateKeysFromState(state *State) []*PrivateKey {
	return state.PrivateKeys
}

func GetServersFromState(state *State) []*Server {
	return state.Servers
}

func GetProjectsFromState(state *State) []*Project {
	return state.Projects
}

func GetApplicationsFromState(state *State) []*Application {
	return state.Applications
}

func GetDatabasesFromState(state *State) []*Database {
	return state.Databases
}
