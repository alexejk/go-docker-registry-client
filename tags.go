package registry

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (r *RegistryClient) Tags(repository string) (tags []string, err error) {

	url := r.url("/v2/%s/tags/list", repository)

	var response tagsResponse
	for {

		r.Log.Debugf("registry.tags url=%s repository=%s", url, repository)
		url, err = r.getPaginatedJson(url, &response)

		switch err {
		case ErrNoMorePages:
			tags = append(tags, response.Tags...)
			return tags, nil
		case nil:
			tags = append(tags, response.Tags...)
			continue
		default:
			return nil, err
		}
	}
}
