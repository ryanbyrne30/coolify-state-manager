package coolify

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type Resource interface {
	GetUUID() string
	SetUUID(string)
	GetIdentifier() string
	ToSaveState() (map[string]interface{}, error)
	ToCreatePayload() ([]byte, error)
	ToUpdatePayload() ([]byte, error)
	BuildNewFromCurrentState(*State) Resource
}

func resourcesToState(resources []Resource) ([]interface{}, error) {
	result := make([]interface{}, len(resources))
	for i, coolify := range resources {
		state, err := coolify.ToSaveState()
		if err != nil {
			return nil, err
		}
		result[i] = state
	}
	return result, nil
}

func structToMap(s interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func structToPayload(s interface{}, ignore_attrs []string) ([]byte, error) {
	m, err := structToMap(s)
	if err != nil {
		return nil, err
	}

	for _, attr := range ignore_attrs {
		delete(m, attr)
	}

	return json.Marshal(m)
}

func hashString(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func hashValues(m map[string]interface{}, attrs []string) {
	for _, a := range attrs {
		s, ok := m[a].(string)
		if ok {
			m[a] = hashString(s)
		}
	}
}

func MergeStates(s1 *State, s2 *State) {
	s1.PrivateKeys = append(s1.PrivateKeys, s2.PrivateKeys...)
	s1.Servers = append(s1.Servers, s2.Servers...)
	s1.Projects = append(s1.Projects, s2.Projects...)
	s1.Applications = append(s1.Applications, s2.Applications...)
	s1.Databases = append(s1.Databases, s2.Databases...)
}
