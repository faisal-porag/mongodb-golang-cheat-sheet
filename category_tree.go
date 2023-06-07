package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	CategoryID int    `bson:"categoryId"`
	ParentID   int    `bson:"parentId"`
	Name       string `bson:"name"`
	Status     string `bson:"status"`
	Children   []*Category // Not used in db. used for response
}

func BulkInsertCategories(categories []Category) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	collection := client.Database("testdb").Collection("category")

	var documents []interface{}
	for _, category := range categories {
		documents = append(documents, category)
	}

	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		return fmt.Errorf("failed to insert documents: %v", err)
	}

	log.Println("Bulk insert completed successfully")

	return nil
}

func bInsert() {
	categories := []Category{
		{CategoryID: 1, ParentID: 0, Name: "Men's Style", Status: "Active"},
		{CategoryID: 2, ParentID: 1, Name: "Shirt", Status: "Active"},
		{CategoryID: 3, ParentID: 1, Name: "Pant", Status: "Active"},
		{CategoryID: 4, ParentID: 2, Name: "Casual", Status: "Active"},
		{CategoryID: 5, ParentID: 2, Name: "Formal", Status: "Active"},
		{CategoryID: 6, ParentID: 3, Name: "Jeans", Status: "Active"},
		{CategoryID: 7, ParentID: 3, Name: "Formal", Status: "Active"},
		{CategoryID: 8, ParentID: 2, Name: "Office", Status: "Active"},
	}

	err := BulkInsertCategories(categories)
	if err != nil {
		log.Fatalf("Failed to perform bulk insert: %v", err)
	}

}

func GenerateCategoryTree(categories []*Category) []*Category {
	// Create a map to store categories by ID
	categoryMap := make(map[int]*Category)

	// Traverse the categories to populate the map
	for _, category := range categories {
		categoryMap[category.CategoryID] = category
	}

	// Create a slice to store root categories
	var roots []*Category

	// Traverse the categories to build the tree
	for _, category := range categories {
		parentID := category.ParentID

		// If the category has a parent, add it as a child to the parent category
		if parent, ok := categoryMap[parentID]; ok {
			parent.Children = append(parent.Children, category)
		} else {
			// If the category has no parent, it is a root category
			roots = append(roots, category)
		}
	}

	return roots
}

func GenerateCategoryTreeV1(categories []*Category, depth int) []*Category {
	// Create a map to store categories by ID
	categoryMap := make(map[int]*Category)

	// Traverse the categories to populate the map
	for _, category := range categories {
		categoryMap[category.CategoryID] = category
	}

	// Create a slice to store root categories
	var roots []*Category

	// Traverse the categories to build the tree
	for _, category := range categories {
		parentID := category.ParentID

		// If the category has a parent, add it as a child to the parent category
		if parent, ok := categoryMap[parentID]; ok {
			if depth > 1 {
				parent.Children = append(parent.Children, category)
			}
		} else {
			// If the category has no parent, it is a root category
			roots = append(roots, category)
		}
	}

	return roots
}

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getAllCategories() ([]*Category, error) {
	client, err := connectDB()
	if err != nil {
		return nil, err
	}

	collection := client.Database("testdb").Collection("category")

	filter := bson.M{} // You can add additional filters if needed

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var categories []*Category
	if err := cur.All(context.Background(), &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func main() {

	categories, _ := getAllCategories()
	// Generate the category tree
	categoryTree := GenerateCategoryTree(categories)

	// Print the category tree in JSON format
	jsonData, _ := json.MarshalIndent(categoryTree, "", "  ")
	fmt.Println(string(jsonData))
}
