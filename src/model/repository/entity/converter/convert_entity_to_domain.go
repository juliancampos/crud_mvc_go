package converter

import (
	"github.com/juliancampos/crud_mvc_go/src/model"
	"github.com/juliancampos/crud_mvc_go/src/model/repository/entity"
)

func ConvertEntityDoDomain(
	entity *entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())
	return domain
}
