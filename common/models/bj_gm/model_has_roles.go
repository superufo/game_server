package bj_gm

type ModelHasRoles struct {
	RoleId    uint64 `xorm:"not null pk UNSIGNED BIGINT"`
	ModelType string `xorm:"not null pk index(model_has_roles_model_id_model_type_index) VARCHAR(255)"`
	ModelId   uint64 `xorm:"not null pk index(model_has_roles_model_id_model_type_index) UNSIGNED BIGINT"`
}

func (m *ModelHasRoles) TableName() string {
	return "model_has_roles"
}
