package sqlp

import (
	"context"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type ProductSQL struct { // Наша структура в которой хранится информация о продукте
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	Created_at  time.Time `json:"time"`
}

func validateName(name string) error { // Валидация имени
	if name == "" || len([]rune(name)) > 20 {
		return WrongName
	}
	return nil
}

func validateDescription(description string) error { // Валидация описания
	if len([]rune(description)) > 300 {
		return CharacterLimitReached
	}
	return nil
}

func validatePrice(price int) error { // Валидация цены
	if price <= 0 {
		return WrongPrice
	}
	return nil
}

func validateQuantity(quantity int) error { // Валидация количества
	if quantity <= 0 {
		return WrongQuantity
	}
	return nil
}

func ValidateALL(name string, description string, price int, quantity int) error { // Валидация всех поступаемых значений
	if err := validateName(name); err != nil {
		return err
	}
	if err := validateDescription(description); err != nil {
		return err
	}
	if err := validatePrice(price); err != nil {
		return err
	}
	if err := validateQuantity(quantity); err != nil {
		return err
	}
	return nil
}

func ValidateKeyValue(key string, value string) error {
	switch {
	case key == "name":
		err := validateName(value)
		return err
	case key == "description":
		err := validateDescription(value)
		return err
	case key == "price":
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		err = validatePrice(v)
		return err
	case key == "quantity":
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		err = validateQuantity(v)
		return err
	default:
		return WrongKey
	}
}

func CheckIdExists(id int, conn *pgx.Conn, ctx context.Context) (bool, error) {
	check := `SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)`
	var exists bool
	err := conn.QueryRow(ctx, check, id).Scan(&exists)
	return exists, err
}
