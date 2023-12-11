package config

func AuthServiceHost() string {
	return "http://localhost:8001"
}

func MasterServiceHost() string {
	return "http://localhost:8002/api"
}

func ProfileServiceHost() string {
	return "http://localhost:8003/api"
}

func MateriServiceHost() string {
	return "http://localhost:8004/api"
}
