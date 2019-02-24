package project

// GameProject request model
type ProjectRequest struct {
	Name            string   `json:"name"`
	BriefSynopsis   string   `json:"value"`
	FullDescription string   `json:"briefSynopsis"`
	Contributors    []string `json:"contributors"`
}
