package mysql

import (
	core "LifeGuard/src/core/db"
	"LifeGuard/src/mpu6050/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(data *entities.Mpu6050) error {
	query := `INSERT INTO mpu6050 (pasos, distancia, fecha, macaddress) VALUES (?, ?, ?, ?)`
	_, err := mysql.conn.DB.Exec(query, data.Pasos, data.Distancia, data.Timestamp, data.MacAddress)
	if err != nil {
		log.Println("Error insertando los valores:", err)
		return err
	}
	fmt.Println("Nuevo dato almacenado con ID:")
	return nil
}

func (mysql *MySQL) GetAll() ([]*entities.Mpu6050, error) {
	query := `SELECT id, pasos, distancia, fecha, macaddress FROM mpu6050`
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("Error consultando datos:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Mpu6050
	for rows.Next() {
		var product entities.Mpu6050
		err := rows.Scan(&product.ID, &product.Pasos, &product.Distancia, &product.Timestamp, &product.MacAddress)
		if err != nil {
			log.Println("Error leyendo fila:", err)
			continue
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error durante la iteración de las filas:", err)
		return nil, err
	}
	if len(products) == 0 {
		return []*entities.Mpu6050{}, nil
	}

	return products, nil
}
