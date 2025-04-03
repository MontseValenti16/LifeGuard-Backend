package mysql

import (
	core "LifeGuard/src/core/db"
	"LifeGuard/src/ds18b20/domain/entities"
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

func (mysql *MySQL) Save(data *entities.Ds18b20) error {
	query := `INSERT INTO ds18b20 (temperatura, fecha, macaddress) VALUES (?, ?, ?)`
	_, err := mysql.conn.DB.Exec(query, data.Temperatura, data.Timestamp, data.MacAddress)
	if err != nil {
		log.Println("Error insertando los valores:", err)
		return err
	}
	fmt.Println("Nuevo dato almacenado con ID:")
	return nil
}

func (mysql *MySQL) GetAll() ([]*entities.Ds18b20, error) {
	query := `SELECT id, temperatura, fecha, macaddress FROM ds18b20`
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("Error consultando datos:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Ds18b20
	for rows.Next() {
		var product entities.Ds18b20
		err := rows.Scan(&product.ID, &product.Temperatura, &product.MacAddress, &product.Timestamp)
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
		return []*entities.Ds18b20{}, nil
	}

	return products, nil
}
