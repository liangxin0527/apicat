package request

import (
	protobase "apicat-cloud/backend/route/proto/base"
	projectbase "apicat-cloud/backend/route/proto/project/base"
)

type GetDefinitionResponseOption struct {
	protobase.ProjectIdOption
	ResponseID uint `uri:"responseID" json:"responseID" query:"responseID" binding:"required,numeric,gt=0"`
}

type CreateDefinitionResponseOption struct {
	protobase.ProjectIdOption
	projectbase.DefinitionResponseDataOption
	projectbase.DefinitionResponseParentIDOption
	projectbase.DefinitionResponseTypeOption
}

type UpdateDefinitionResponseOption struct {
	GetDefinitionResponseOption
	projectbase.DefinitionResponseDataOption
}

type DeleteDefinitionResponseOption struct {
	GetDefinitionResponseOption
	projectbase.DerefOption
}

type SortDefinitionResponseOption struct {
	SortOption
}
