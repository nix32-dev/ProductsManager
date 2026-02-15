package sqlp

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

var CTX context.Context = context.Background()
var Conn *pgx.Conn

func CreateConnection(ctx context.Context, connection string) (*pgx.Conn, error) {
	connect, err := pgx.Connect(ctx, connection) // Подключаемся к базе данных
	if err != nil {
		return nil, err
	}

	fmt.Println("Подключение к PostgreSQL произошло успешно! | The connection to PostgreSQL was successful!") // Индикация в консоль, что подключение прошло успешно
	return connect, nil
}

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlRequest := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		description VARCHAR(300),
		price INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	/* создаем таблицу с столбцами:
	id (уникальный индетификатор товаров),
	name (название - до 20 символов),
	description(описание - до 300 символов),
	price (цена товара),
	quantity (количество товара),
	created_at (время создания товара как единицу)  */

	_, err := conn.Exec(ctx, sqlRequest) // Отправляем запрос на создание
	return err
}

func CreateProduct(ctx context.Context, conn *pgx.Conn, name string, description string, price int, quantity int) (ProductSQL, error) {
	// Описываем наш SQL запрос
	sqlRequest := `
        INSERT INTO products 
        (name, description, price, quantity, created_at) 
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := conn.Exec(ctx, sqlRequest, name, description, price, quantity, time.Now().Format(time.RFC3339))              // Отправляем полученные данные в базу данных (После валидации)
	return ProductSQL{Name: name, Description: description, Price: price, Quantity: quantity, Created_at: time.Now()}, err // Возвращаем ошибку при наличии и созданный продукт
}

func ChangeProduct(ctx context.Context, conn *pgx.Conn, id int, key string, value string) error {
	var v int
	var sqlRequest string
	var err error

	switch { // Смотрим на значения, проверяем, меняем.

	case key == "price": // Если price, переводим в int, делаем валидацию, отгружаем в SQL
		v, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
		if err = validatePrice(v); err != nil {
			return err
		}
		sqlRequest = `
		UPDATE products
		SET price = $1
		WHERE id = $2`

		_, err := conn.Exec(ctx, sqlRequest, v, id)
		return err
	case key == "quantity": // Если quantity, переводим в int, делаем валидацию, отгружаем в SQL
		v, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
		if err = validateQuantity(v); err != nil {
			return err
		}
		sqlRequest = `
		UPDATE products
		SET quantity = $1
		WHERE id = $2`

		_, err := conn.Exec(ctx, sqlRequest, v, id)
		return err
	case key == "name": // Если name, делаем валидацию, отгружаем в SQL
		if err = validateName(value); err != nil {
			return err
		}
		sqlRequest = `
		UPDATE products
		SET name = $1
		WHERE id = $2`

		_, err := conn.Exec(ctx, sqlRequest, value, id)
		return err
	case key == "description": // Если name, делаем валидацию, отгружаем в SQL
		if err = validateDescription(value); err != nil {
			return err
		}
		sqlRequest = `
		UPDATE products
		SET description = $1
		WHERE id = $2`

		_, err := conn.Exec(ctx, sqlRequest, value, id)
		return err
	default:
		return WrongKey
	}
}

func DeleteProduct(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlRequest := `
	DELETE FROM products
	WHERE id = $1`
	_, err := conn.Exec(ctx, sqlRequest, id) // Создаем SQL запрос и удаляем продукт по его ID
	return err
}

func GetProduct(ctx context.Context, conn *pgx.Conn, pageR string) ([]ProductSQL, error) {
	page, err := strconv.Atoi(pageR) // Переводим номер страницы в int, проверяем
	if err != nil {
		return nil, err
	}
	page -= 1
	page *= 10
	sqlRequest := `
	SELECT id, name, description, price, quantity, created_at FROM products
	ORDER BY id
	LIMIT 10
	OFFSET $1`

	getProduct, err := conn.Query(ctx, sqlRequest, page)
	var products []ProductSQL
	for getProduct.Next() { // Делаем слайс и подгружаем данные с таблицы
		var product ProductSQL
		if err = getProduct.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Created_at); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	getProduct.Close()
	return products, nil
}
