package entities

type Ds18b20 struct{
	ID 		  int `json:"id"`
	Temperatura float64 `json:"temperatura"`
	Timestamp string    `json:"fecha"`
	MacAddress string   `json:"macaddress"`
	IdPersona int `json:"id_persona"`
}