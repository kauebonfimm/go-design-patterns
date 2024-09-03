package adapter

import "encoding/json"

type ShopifyAdapter struct {
}

func (a *ShopifyAdapter) GetCategories(id string) []Category {
	var categories []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Details struct {
			Slug        string `json:"slug"`
			Description string `json:"description"`
		} `json:"details"`
	}

	data := `[{"id":1,"name":"1","details":{"description":"1","slug":"1"}},{"id":2,"name":"2","details":{"description":"2","slug":"2"}},{"id":3,"name":"3","details":{"description":"3","slug":"3"}},{"id":4,"name":"4","details":{"description":"4","slug":"4"}},{"id":6,"name":"6","details":{"description":"6","slug":"6"}}]`

	json.Unmarshal([]byte(data), &categories)

	res := make([]Category, 0)

	for _, v := range categories {
		res = append(res, Category{
			Name:        v.Name,
			Description: v.Details.Description,
			Slug:        v.Details.Slug,
		})
	}

	return res
}
