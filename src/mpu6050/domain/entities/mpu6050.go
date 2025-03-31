package entities

type Mpu6050 struct{
	ID 		  int `json:"id"`
	Pasos	  int `json:"pasos"`
	Distancia int `json:"distancia"`
}