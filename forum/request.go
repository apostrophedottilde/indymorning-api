package forum

// Forum request model
type Request struct {
	Title       string   `json:"title" bson:"title,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Posts       []Post   `json:"posts" bson:"posts,omitempty"`
	Creator     string   `json:"creator" bson:"creator,omitempty"`
	Tags        []string `json:"tags" bson:"tags,omitempty"`
	State       string   `json:"state" bson:"state,omitempty"`
}
