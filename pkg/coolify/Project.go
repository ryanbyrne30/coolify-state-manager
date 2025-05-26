package coolify

type Project struct {
	Id          string `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p *Project) GetUUID() string {
	return p.UUID
}

func (p *Project) SetUUID(uuid string) {
	p.UUID = uuid
}

func (p *Project) GetIdentifier() string {
	return p.Id
}

func (p *Project) ToSaveState() (map[string]interface{}, error) {
	m, err := structToMap(p)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (p *Project) ToCreatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid"}
	return structToPayload(p, ignore_attrs)
}

func (p *Project) ToUpdatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid"}
	return structToPayload(p, ignore_attrs)
}

func (p *Project) BuildNewFromCurrentState(state *State) Resource {
	copy := *p
	return &copy
}
