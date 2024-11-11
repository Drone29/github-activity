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
		Ref     string `json:"ref,omitempty"`
		RefType string `json:"ref_type,omitempty"` // fo CreateEvent
		Issue   struct {
			State string `json:"state,omitempty"`
		} `json:"issue,omitempty"` // IssuesEvent
	} `json:"payload"`
}

type Activities struct {
	Objects []Activity
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
	for i, a := range as.Objects {
		obj_s := a.String()
		result += obj_s
		if len(obj_s) > 0 && i < len(as.Objects)-1 {
			result += ",\n"
		}
	}
	return result + "\n]"
}

func ParseActivitiesAsArray(data []byte) ([]Activity, error) {
	activities := make([]Activity, 0)
	err := json.Unmarshal(data, &activities)
	return activities, err
}

func ParseActivities(data []byte) (Activities, error) {
	activities, err := ParseActivitiesAsArray(data)
	return Activities{activities}, err
}
