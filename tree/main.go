package main

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	ID       string     `json:"ID"`
	ParentID string     `json:"ParentID"`
	Name     string     `json:"Name"`
	Status   string     `json:"Status"`
	Children []Category `json:"Children"`
}

func hasChildren(category Category) bool {
	return len(category.Children) > 0
}

func getNestedIDs(category Category) []string {
	var nestedIDs []string
	for _, child := range category.Children {
		if hasChildren(child) {
			nestedIDs = append(nestedIDs, child.ID)
			nestedIDs = append(nestedIDs, getNestedIDs(child)...)
		}
	}
	return nestedIDs
}

func main() {
	data := `[{"ID":"1","ParentID":"","Name":"Men","Status":"Active","Children":[{"ID":"2","ParentID":"1","Name":"Shirts","Status":"Active","Children":[{"ID":"4","ParentID":"2","Name":"Formal Shirts","Status":"Active","Children":[]},{"ID":"5","ParentID":"2","Name":"Casual Shirts","Status":"Active","Children":[{"ID":"8","ParentID":"5","Name":"Formal Shirts","Status":"Active","Children":[]},{"ID":"9","ParentID":"5","Name":"Casual Shirts","Status":"Active","Children":[{"ID":"10","ParentID":"9","Name":"Formal Shirts","Status":"Active","Children":[]}]}]}]},{"ID":"3","ParentID":"1","Name":"Pants","Status":"Active","Children":[{"ID":"6","ParentID":"3","Name":"Jeans","Status":"Active","Children":[]},{"ID":"7","ParentID":"3","Name":"Chinos","Status":"Active","Children":[]}]}]}]`

	var categories []Category
	err := json.Unmarshal([]byte(data), &categories)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var nestedIDs []string
	for _, category := range categories {
		if hasChildren(category) {
			nestedIDs = append(nestedIDs, getNestedIDs(category)...)
		}
	}

	fmt.Println("Nested IDs:", nestedIDs)
}
