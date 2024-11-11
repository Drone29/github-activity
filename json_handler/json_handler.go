package json_handler

import (
	"encoding/json"
)

type Activity struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Public    bool   `json:"public"`
	Repo      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Action  string `json:"action,omitempty"`
		Commits []struct {
			Author struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"commits,omitempty"`
	} `json:"payload"`
}

type Activities struct {
	Activities []Activity
}

// stringer for individual activities
func (a Activity) String() string {
	data, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return ""
	}
	return string(data)
}

// stringer for a slice of activities
func (as Activities) String() string {
	var result string = "[\n"
	for i, a := range as.Activities {
		obj_s := a.String()
		result += obj_s
		if len(obj_s) > 0 && i < len(as.Activities)-1 {
			result += ",\n"
		}
	}
	return result + "\n]"
}

func ParseActivities(data []byte) (Activities, error) {
	activities := make([]Activity, 0)
	err := json.Unmarshal(data, &activities)
	return Activities{activities}, err
}
