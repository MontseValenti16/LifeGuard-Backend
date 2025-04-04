package entities

type Max30102 struct {
    ID        int       `json:"id"`
    BPM       int       `json:"bpm"`
    SpO2      float64   `json:"spo2"`
	Timestamp string    `json:"fecha"`
	MacAddress string   `json:"macaddress"`
    IdPersona int `json:"id_persona"`
}