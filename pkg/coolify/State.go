package coolify

type State struct {
	PrivateKeys  []*PrivateKey  `json:"private_keys"`
	Servers      []*Server      `json:"servers"`
	Projects     []*Project     `json:"projects"`
	Applications []*Application `json:"applications"`
	Databases    []*Database    `json:"databases"`
}

func NewState() *State {
	return &State{
		PrivateKeys:  make([]*PrivateKey, 0),
		Servers:      make([]*Server, 0),
		Projects:     make([]*Project, 0),
		Applications: make([]*Application, 0),
		Databases:    make([]*Database, 0),
	}
}

func (s *State) ToSaveState() (map[string]interface{}, error) {
	x_private_keys := make([]Resource, len(s.PrivateKeys))
	for i, obj := range s.PrivateKeys {
		x_private_keys[i] = obj
	}
	private_keys, err := resourcesToState(x_private_keys)
	if err != nil {
		return nil, err
	}

	x_servers := make([]Resource, len(s.Servers))
	for i, obj := range s.Servers {
		x_servers[i] = obj
	}
	servers, err := resourcesToState(x_servers)
	if err != nil {
		return nil, err
	}

	x_projects := make([]Resource, len(s.Projects))
	for i, obj := range s.Projects {
		x_projects[i] = obj
	}
	projects, err := resourcesToState(x_projects)
	if err != nil {
		return nil, err
	}

	x_applcations := make([]Resource, len(s.Applications))
	for i, obj := range s.Applications {
		x_applcations[i] = obj
	}
	applications, err := resourcesToState(x_applcations)
	if err != nil {
		return nil, err
	}

	x_databases := make([]Resource, len(s.Databases))
	for i, obj := range s.Databases {
		x_databases[i] = obj
	}
	databases, err := resourcesToState(x_databases)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"private_keys": private_keys,
		"servers":      servers,
		"projects":     projects,
		"applications": applications,
		"databases":    databases,
	}, nil
}

func (s *State) serverUUID(server_id string) string {
	for _, r := range s.Servers {
		if r.Id == server_id {
			return r.UUID
		}
	}
	return ""
}

func (s *State) projectUUID(project_id string) string {
	for _, r := range s.Projects {
		if r.Id == project_id {
			return r.UUID
		}
	}
	return ""
}

func (s *State) privateKeyUUID(private_key_id string) string {
	for _, r := range s.Servers {
		if r.Id == private_key_id {
			return r.UUID
		}
	}
	return ""
}
