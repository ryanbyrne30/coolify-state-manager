package coolify

type Server struct {
	Id             string `json:"id"`
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	IP             string `json:"ip"`
	User           string `json:"user"`
	Port           int    `json:"port"`
	PrivateKeyUUID string `json:"private_key_uuid"`
	PrivateKeyId   string `json:"private_key_id"`
}

func (s *Server) GetUUID() string {
	return s.UUID
}

func (s *Server) SetUUID(uuid string) {
	s.UUID = uuid
}

func (s *Server) GetIdentifier() string {
	return s.Id
}

func (s *Server) ToSaveState() (map[string]interface{}, error) {
	m, err := structToMap(s)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Server) ToCreatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "private_key_id"}
	return structToPayload(s, ignore_attrs)
}

func (s *Server) ToUpdatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "private_key_id", "private_key_uuid"}
	return structToPayload(s, ignore_attrs)
}

func (s *Server) BuildNewFromCurrentState(state *State) Resource {
	copy := *s
	copy.PrivateKeyUUID = state.privateKeyUUID(copy.PrivateKeyId)
	return &copy
}
