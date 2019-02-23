package project

// GameProject request model
type GameProjectRequest struct {
	Name            string   `json:"name"`
	BriefSynopsis   string   `json:"value"`
	FullDescription string   `json:"briefSynopsis"`
	Contributors    []string `json:"contributors"`
}
