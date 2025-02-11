package middlewares

import (
	_ "fmt"
	"net/http"

	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	"github.com/PureMLHQ/PureML/packages/purebackend/model/models"
	"github.com/labstack/echo/v4"
)

func ValidateModelBranchVersion(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			modelBranchVersion := context.Param("version")
			modelBranchUUID := context.Get("ModelBranch").(*models.ModelBranchNameResponse).UUID
			if modelBranchVersion == "" {
				context.Response().WriteHeader(http.StatusBadRequest)
				_, err := context.Response().Writer.Write([]byte("Version required"))
				if err != nil {
					return err
				}
				return nil
			}
			version, err := app.Dao().GetModelBranchVersion(modelBranchUUID, modelBranchVersion)
			if err != nil {
				context.Response().WriteHeader(http.StatusInternalServerError)
				_, err = context.Response().Writer.Write([]byte(err.Error()))
				if err != nil {
					return err
				}
				return nil
			}
			if version == nil {
				context.Response().WriteHeader(http.StatusNotFound)
				_, err = context.Response().Writer.Write([]byte("Model Branch Version not found"))
				if err != nil {
					return err
				}
				return nil
			}
			context.Set(ContextModelBranchVersionKey, &models.ModelBranchVersionNameResponse{
				UUID:    version.UUID,
				Version: version.Version,
			})
			return next(context)
		}
	}
}
