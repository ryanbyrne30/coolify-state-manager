package coolify

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ResourceService[T Resource] struct {
	resourceType   string
	endpoint       string
	createEndpoint string
	request        *CoolifyRequestService
}

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func NewResourceService[T Resource](resourceType string, endpoint string, createEndpoint string, request *CoolifyRequestService) *ResourceService[T] {
	return &ResourceService[T]{
		resourceType:   resourceType,
		endpoint:       endpoint,
		createEndpoint: createEndpoint,
		request:        request,
	}
}

func (svc *ResourceService[T]) Create(r T) (string, error) {
	fmt.Printf("Creating coolify: %s [%s]\n", svc.resourceType, r.GetIdentifier())
	data, err := r.ToCreatePayload()
	if err != nil {
		return "", err
	}

	resp, err := svc.request.Post(svc.createEndpoint, data)
	if err != nil {
		return "", err
	}

	var createResponse CreateResponse
	err = json.Unmarshal(resp, &createResponse)
	if err != nil {
		return "", err
	}

	return createResponse.UUID, nil
}

func (svc *ResourceService[T]) Update(r T) error {
	fmt.Printf("Updating coolify: %s [%s]\n", svc.resourceType, r.GetIdentifier())
	data, err := r.ToUpdatePayload()
	if err != nil {
		return err
	}
	_, err = svc.request.Patch(fmt.Sprintf("%s/%s", svc.endpoint, r.GetUUID()), data)
	return err
}

func (svc *ResourceService[T]) Delete(uuid string) error {
	fmt.Printf("Deleting coolify: %s [%s]\n", svc.resourceType, uuid)
	_, err := svc.request.Delete(fmt.Sprintf("%s/%s", svc.endpoint, uuid))
	return err
}

func (svc *ResourceService[T]) SaveState(current_state *State, future_state *State, get_resources func(*State) []T) ([]T, error) {
	// Remove resources that no longer should exist
	current_resources := get_resources(current_state)
	future_resources := get_resources(future_state)
	for _, cur := range current_resources {
		exists := false
		for _, future := range future_resources {
			if cur.GetIdentifier() == future.GetIdentifier() {
				exists = true
				break
			}
		}
		if !exists {
			svc.Delete(cur.GetUUID())
		}
	}

	// Create/Update resources that should exist
	end_result := make([]T, len(future_resources))
	for i, future := range future_resources {
		var existing T
		for _, cur := range current_resources {
			if future.GetIdentifier() == cur.GetIdentifier() {
				existing = cur
			}
		}

		result := future.BuildNewFromCurrentState(current_state)
		if !reflect.ValueOf(existing).IsNil() {
			result.SetUUID(existing.GetUUID())
			svc.Update(result.(T))
		} else {
			uuid, err := svc.Create(result.(T))
			if err != nil {
				return nil, err
			}
			result.SetUUID(uuid)
		}
		end_result[i] = result.(T)
	}
	return end_result, nil
}

func (svc *ResourceService[T]) DestroyState(state *State, get_resources func(*State) []T) error {
	for _, r := range get_resources(state) {
		if err := svc.Delete(r.GetUUID()); err != nil {
			return err
		}
	}
	return nil
}
