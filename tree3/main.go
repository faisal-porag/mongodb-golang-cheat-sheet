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

func findCategoryByID(categories []Category, parentID string) *Category {
	for _, category := range categories {
		if category.ID == parentID {
			return &category
		}
		if category.Children != nil {
			if foundCategory := findCategoryByID(category.Children, parentID); foundCategory != nil {
				return foundCategory
			}
		}
	}
	return nil
}

func hasChildren(category Category) bool {
	return len(category.Children) > 0
}

func getNestedIDsByParentID(categories []Category, parentID string) []string {
	category := findCategoryByID(categories, parentID)
	if category == nil {
		return nil
	}

	var nestedIDs []string
	var queue []Category

	queue = append(queue, *category)

	for len(queue) > 0 {
		currentCategory := queue[0]
		queue = queue[1:]

		for _, child := range currentCategory.Children {
			nestedIDs = append(nestedIDs, child.ID)
			if hasChildren(child) {
				queue = append(queue, child)
			}
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

	parentID := "5"
	nestedIDs := getNestedIDsByParentID(categories, parentID)
	fmt.Println("Nested IDs:", nestedIDs)
}
