package audit

import (
	"context"

	"github.com/Caris-Project/caris-modules/utils"
	"github.com/kr/pretty"
)

// CreatePayload ...
type CreatePayload struct {
	Target   string
	TargetID string
	Action   string
	Data     string
	Author   CreatePayloadAuthor
}

// CreatePayloadAuthor ...
type CreatePayloadAuthor struct {
	ID   string
	Name string
}

// Create ...
func (s Service) Create(payload CreatePayload) {
	ctx := context.Background()

	// Get document
	doc := Audit{
		ID:       utils.GenerateMongoID(),
		Source:   s.Source,
		Target:   payload.Target,
		TargetID: payload.TargetID,
		Action:   payload.Action,
		Data:     payload.Data,
		Author: Author{
			ID:   payload.Author.ID,
			Name: payload.Author.Name,
		},
		CreatedAt: utils.TimeNow(),
	}

	// Insert to db
	colName := getColName(s.Source, payload.Target)
	if _, err := s.DB.Collection(colName).InsertOne(ctx, doc); err != nil {
		pretty.Println("*** Audit create log error", err.Error())
		pretty.Println("*** Payload", payload)
		pretty.Println("*****************")
	}
}
