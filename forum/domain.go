package forum

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type forum interface {
	Close() (forum, error)
	SubmitPost(userID string, forumId string, postText string) (forum, error)
	SubmitComment(userID string, postId string, postText string) (forum, error)
}

type Forum struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Posts       []Post             `json:"posts" bson:"posts,omitempty"`
	Creator     string             `json:"creator" bson:"creator,omitempty"`
	Tags        []string           `json:"tags" bson:"tags,omitempty"`
	State       string             `json:"state" bson:"state,omitempty"`
}

type Post struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Heading  string             `json:"heading" bson:"heading,omitempty"`
	Text     string             `json:"text" bson:"text,omitempty"`
	Comments []Comment          `json:"comments" bson:"comments,omitempty"`
	Creator  string             `json:"creator" bson:"creator,omitempty"`
	Tags     []string           `json:"tags" bson:"tags,omitempty"`
	State    string             `json:"state" bson:"state,omitempty"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text    string             `json:"text" bson:"text,omitempty"`
	Creator string             `json:"creator" bson:"creator,omitempty"`
	Tags    []string           `json:"tags" bson:"tags,omitempty"`
	State   string             `json:"state" bson:"state,omitempty"`
}

func (gp *Forum) Close() (forum, error) {
	if gp.State == "CLOSED" {
		return nil, errors.New("Cannot transition a cancelled project into 'CANCELLED' state")
	}
	gp.State = "CLOSED"
	return gp, nil
}

// SubmitPost to list of contributors to this project
func (gp *Forum) SubmitPost(userID string, forumId string, postText string) (forum, error) {
	gp.Posts = append(gp.Posts, Post{})
	return gp, nil
}

func (gp *Forum) SubmitComment(userID string, forumId string, postText string) (forum, error) {
	gp.Posts = append(gp.Posts, Post{})
	return gp, nil
}
