package project

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Project interface {
	Cancel() (Project, error)
	AddContributor(userID string) (Project, error)
}

type GameProject struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	BriefSynopsis   string             `json:"value" bson:"value,omitempty"`
	FullDescription string             `json:"briefSynopsis" bson:"briefSynopsis,omitempty"`
	Creator         string             `json:"creator" bson:"creator,omitempty"`
	Contributors    []string           `json:"contributors" bson:"contributors,omitempty"`
	State           string             `json:"state" bson:"state,omitempty"`
}

// Cancel a project - doesn't delete but sets state to 'CANCELLED'
func (gp *GameProject) Cancel() (Project, error) {
	if gp.State == "CANCELLED" {
		return nil, errors.New("Cannot transition a cancelled project into 'CANCELLED' state")
	}
	gp.State = "CANCELLED"
	return gp, nil
}

// AddContributor to list of contributors to this project
func (gp *GameProject) AddContributor(userID string) (Project, error) {
	gp.Contributors = append(gp.Contributors, userID)
	return gp, nil
}
