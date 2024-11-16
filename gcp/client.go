package gcp

func CreateClient(Project string) *Client {
	client := &Client{Project: Project}
	return client
}
