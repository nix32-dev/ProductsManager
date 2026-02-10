package sqlp

import (
	"errors"
)

var WrongRequestMethod error = errors.New("Ошибка: Неверный метод подключения!\nError: Incorrect connection method!")
var WrongName error = errors.New("Указано неверное имя продукта!\nThe product name is incorrect!")
var CharacterLimitReached error = errors.New("Превышен лимит символов!\nThe character limit has been exceeded!")
var WrongPrice error = errors.New("Цена указана неверно!\nThe price is incorrect!")
var WrongQuantity error = errors.New("Количество указано неверно!\nThe quantity is incorrect!")
var WrongKey error = errors.New("Неверно указан ключ таблицы!\nThe table key is specified incorrectly!")
var NotExists error = errors.New("Продукта по указанному ID не существует!\nThe product with the specified ID does not exist!")
var WrongPage error = errors.New("Страница указана неверно!\nThe page is specified incorrectly!")