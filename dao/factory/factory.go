package factory

import (
	"../interfaces"
	"../mysqld"
)

func FactoryDao(e string) interfaces.UserDao {
	var i interfaces.UserDao = mysqld.UserImplMysql{}
	return i
}
