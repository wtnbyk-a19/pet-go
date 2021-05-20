package api

import (
	"pet-go/middlewares"

	"github.com/labstack/echo"
)

func RegisterPhysique() echo.HandlerFunc {
	return func(c echo.Context) error {

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		// petの検索
		// 例外処理

		// モデルの登録処理

	}
}
