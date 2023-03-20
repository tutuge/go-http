package main

type Result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Kubernetes struct {
					NamespaceLabels struct {
						Name string `json:"name"`
					} `json:"namespace_labels"`
					Deployment struct {
						Name string `json:"name"`
					} `json:"deployment"`
				} `json:"kubernetes"`
				Message string `json:"message"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
