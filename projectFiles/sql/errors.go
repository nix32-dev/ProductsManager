package sqlp

import (
	"errors"
)

var WrongRequestMethod error = errors.New("Неверный метод подключения! | Incorrect connection method!")
var WrongName error = errors.New("Указано неверное имя продукта! | The product name is incorrect!")
var CharacterLimitReached error = errors.New("Превышен лимит символов! | The character limit has been exceeded!")
var WrongPrice error = errors.New("Цена указана неверно! | The price is incorrect!")
var WrongQuantity error = errors.New("Количество указано неверно! | The quantity is incorrect!")
var WrongKey error = errors.New("Неверно указан ключ таблицы! | The table key is specified incorrectly!")
var NotExists error = errors.New("Продукта по указанному ID не существует! | The product with the specified ID does not exist!")
var WrongPage error = errors.New("Страница указана неверно! | The page is specified incorrectly!")
