package coolify

type Database struct {
	Id               string `json:"id"`
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ServerId         string `json:"server_id"`
	ServerUUID       string `json:"server_uuid"`
	ProjectId        string `json:"project_id"`
	ProjectUUID      string `json:"project_uuid"`
	EnvironmentName  string `json:"environment_name"`
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDB       string `json:"postgres_db"`
	Image            string `json:"image"`
	LimitsCpus       string `json:"limits_cpus"`
	LimitsMemory     string `json:"limits_memory"`
}

func (db *Database) GetUUID() string {
	return db.Id
}

func (db *Database) SetUUID(uuid string) {
	db.UUID = uuid
}

func (db *Database) GetIdentifier() string {
	return db.Id
}

func (db *Database) ToSaveState() (map[string]interface{}, error) {
	m, err := structToMap(db)
	if err != nil {
		return nil, err
	}

	hashValues(m, []string{
		"postgres_password",
	})

	return m, nil
}

func (db *Database) ToCreatePayload() ([]byte, error) {
	ignore_attrs := []string{"id", "uuid", "server_id", "project_id"}
	return structToPayload(db, ignore_attrs)
}

func (db *Database) ToUpdatePayload() ([]byte, error) {
	ignore_attrs := []string{
		"id",
		"uuid",
		"project_id",
		"server_id",
		"project_uuid",
		"server_uuid",
		"environment_name",
	}
	return structToPayload(db, ignore_attrs)
}

func (db *Database) BuildNewFromCurrentState(state *State) Resource {
	copy := *db
	copy.ServerUUID = state.serverUUID(db.ServerId)
	copy.ProjectUUID = state.projectUUID(db.ProjectId)
	return &copy
}
