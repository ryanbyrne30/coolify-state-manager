package coolify

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type CoolifyService struct {
	current_state        *State
	servers_service      *ResourceService[*Server]
	private_keys_service *ResourceService[*PrivateKey]
	projects_service     *ResourceService[*Project]
	applications_service *ResourceService[*Application]
	databases_service    *ResourceService[*Database]
}

func NewCoolifyService(api_url string, token string, current_state *State) *CoolifyService {
	request := NewCoolifyRequestService(api_url, token)
	return &CoolifyService{
		current_state:        current_state,
		servers_service:      NewResourceService[*Server]("Server", "/api/v1/servers", "/api/v1/servers", request),
		applications_service: NewResourceService[*Application]("Application", "/api/v1/applications", "/api/v1/applications/private-deploy-key", request),
		databases_service:    NewResourceService[*Database]("Database", "/api/v1/databases", "/api/v1/databases/postgresql", request),
		projects_service:     NewResourceService[*Project]("Project", "/api/v1/projects", "/api/v1/projects", request),
		private_keys_service: NewResourceService[*PrivateKey]("PrivateKey", "/api/v1/security/keys", "/api/v1/security/keys", request),
	}
}

func (svc *CoolifyService) Apply(new_state *State, state_file string) {
	updated_state, err := svc.private_keys_service.SaveState(svc.current_state, new_state, GetPrivateKeysFromState)
	if err != nil {
		fmt.Printf("Error saving private keys state: %v", err)
		panic(err)
	}
	svc.current_state.PrivateKeys = updated_state
	svc.SaveState(state_file)
}

func (svc *CoolifyService) Destroy(state_file string) error {
	err := svc.private_keys_service.DestroyState(svc.current_state, GetPrivateKeysFromState)
	if err != nil {
		return err
	}
	svc.current_state.PrivateKeys = []*PrivateKey{}
	svc.SaveState(state_file)

	return nil
}

func (svc *CoolifyService) SaveState(ofile string) {
	file, err := os.Create(ofile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	state, err := svc.current_state.ToSaveState()
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(state)
	if err != nil {
		panic(err)
	}
}

func ReadState(ifile string) (*State, error) {
	file, err := os.Open(ifile)
	if err != nil {
		return NewState(), nil
	}
	defer file.Close()

	var state State
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		if err == io.EOF {
			return NewState(), nil
		}
		return nil, err
	}

	return &state, nil
}
