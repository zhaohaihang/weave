package model

type Group struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:100;not null;unique"`
	Describe  string `json:"describe" gorm:"size:1024;"`
	CreatorId uint   `json:"creatorId"`
	UpdaterId uint   `json:"updaterId"`
	Users     []User `json:"users" gorm:"many2many:user_groups;"`

	BaseModel
}

type CreatedGroup struct {
	Name      string `json:"name"`
	Describe  string `json:"describe"`
	CreatorId uint   `json:"creatorId"`
}

func (g *CreatedGroup) GetGroup(uid uint) *Group {
	return &Group{
		Name:      g.Name,
		Describe:  g.Describe,
		CreatorId: g.CreatorId,
	}
}

type UpdatedGroup struct {
	Name      string `json:"name"`
	Describe  string `json:"describe"`
	UpdaterId uint   `json:"updaterId"`
}

func (g *UpdatedGroup) GetGroup(uid uint) *Group {
	return &Group{
		Name:      g.Name,
		Describe:  g.Describe,
		UpdaterId: g.UpdaterId,
	}
}
