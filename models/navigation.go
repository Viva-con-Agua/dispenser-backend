package models

type (
	Label struct {
		DE_DE string `json:"de_DE" validator:"required"`
		EN_EN string `json:"en_EN" validator:"required"`
	}
	Entry struct {
		Label    Label   `json:"label" validator:"required"`
		Url      string  `json:"url" validator:"required"`
		Entries  Entries `json:"entries"`
		RoleName string  `json:"role_name" validator:"required"`
		Style    string  `json:"style" validator:"required"`
	}
	Entries    []Entry
	Navigation struct {
		Name    string  `json:"name" validator:"required"`
		Entries Entries `json:"entries" validator:"required"`
	}
)

func (e *Entries) Restrict() *Entries {
	return e

}
