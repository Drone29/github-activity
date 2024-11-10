package main

import (
	"encoding/json"
	"fmt"
	"github-activity/http_handler"
	"log"
	"os"
	"strconv"
	"strings"
)

type Activity struct {
	Type      string
	Public    string
	ID        string
	CreatedAt string
}

func find_values(data any, keys map[string]struct{}, result map[string]string) bool {

	// iterate over keys
	switch decoded := data.(type) {
	case map[string]interface{}:
		// json object
		for k, v := range decoded {
			if _, ok := keys[k]; ok {
				switch value := v.(type) {
				case string:
					result[k] = value
				case int:
					result[k] = strconv.Itoa(value)
				case bool:
					result[k] = strconv.FormatBool(value)
				}

				delete(keys, k)
				if len(keys) == 0 {
					return true
				}
			}

			// look for nested structures
			switch nested := v.(type) {
			case map[string]interface{}, []interface{}:
				// check composite key
				for comp_k := range keys {
					strtok := strings.Split(comp_k, ".")
					if len(strtok) > 1 && strtok[0] == k {
						// look for shortened key
						keys_nested := make(map[string]struct{})
						res_nested := make(map[string]string)
						new_key := strings.Join(strtok[1:], ".")
						keys_nested[new_key] = struct{}{}
						if find_values(nested, keys_nested, res_nested) {
							result[comp_k] = res_nested[new_key]

							delete(keys, comp_k)
							if len(keys) == 0 {
								return true
							}
						}
						break
					}
				}
				if find_values(nested, keys, result) {
					return true
				}
			}
		}
	case []interface{}:
		// json array
		for _, val := range decoded {
			if find_values(val, keys, result) {
				return true
			}
		}
	}

	return len(keys) == 0
}

func FindValues(data any, keys []string) (values map[string]string, ok bool) {
	keys_map := make(map[string]struct{})
	values = make(map[string]string)
	// convert slice to map
	for _, k := range keys {
		keys_map[k] = struct{}{}
	}
	ok = find_values(data, keys_map, values)
	return values, ok
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments!")
		fmt.Println(os.Args[0], "<username>")
		return
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events", os.Args[1])
	body, _, err := http_handler.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch %v\n", err)
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to decode json %v\n", err)
	}

	vals, ok := FindValues(data, []string{"type", "public", "id", "created_at", "repo.name"})
	if !ok {
		log.Fatalf("Couldn't find all the values\n")
	}
	fmt.Println(vals)

}
