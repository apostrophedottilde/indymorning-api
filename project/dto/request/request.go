package request

// GameProject request model
type GameProject struct {
	Name            string   `json:"name"`
	BriefSynopsis   string   `json:"value"`
	FullDescription string   `json:"briefSynopsis"`
	Contributors    []string `json:"contributors"`
}
