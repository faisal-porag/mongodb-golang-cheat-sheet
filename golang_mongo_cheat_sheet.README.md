


> document that have atleast one product with name apple and quantity 15 in json array using mongo in go 

```go
type Product struct {
	Name     string `bson:"name"`
	Quantity int    `bson:"quantity"`
}

type Document struct {
	ID       int       `bson:"_id"`
	Products []Product `bson:"products"`
}
```

```go
// Define the query
	filter := bson.M{
		"products": bson.M{
			"$elemMatch": bson.M{
				"name":     "apple",
				"quantity": bson.M{"$gt": 15},
			},
		},
	}
```

> get only matched item from json array 

```go
// Define the query
	filter := bson.M{
		"products": bson.M{
			"$elemMatch": bson.M{
				"name":     "apple",
				"quantity": bson.M{"$gt": 15},
			},
		},
	}

	// Define the projection to return only the matched products
	projection := bson.M{
		"products.$": 1,
	}
```

---

> `like` in mongo using go

```go
// Define the search pattern
	searchPattern := "app" // Example search pattern, e.g., "app" matches "apple", "application", etc.

	// Create a regular expression pattern for the "name" field
	regexPattern := bson.M{
		"$regex":   searchPattern,
		"$options": "i", // Case-insensitive search
	}

	// Define the filter query
	filter := bson.M{
		"name": regexPattern,
	}
```

---

> Json array size count in mongo using go 

```go
// Define the filter query
	filter := bson.M{
		"products": bson.M{
			"$exists": true, // Optional: Include this if you want to filter documents that have the "products" field
			"$size":   3,    // Specify the desired array size
		},
	}

	// Execute the count query
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Array size count:", count)
```

---


> Compair with two field in mongo using go

```go
// Define the filter query with $expr
	filter := bson.M{
		"$expr": bson.M{
			"$gt": []interface{}{
				"$quantity",
				"$price",
			},
		},
	}
```

---


> Multiple collection join 

```go
// Get a reference to the MongoDB database and collections
	db := client.Database("your_database")
	productCollection := db.Collection("product")
	brandsCollection := db.Collection("brands")
	categoryCollection := db.Collection("category")
	variantsCollection := db.Collection("product_variants")

	// Define your pipeline stages
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         brandsCollection.Name(),
				"localField":   "brand_id",
				"foreignField": "_id",
				"as":           "brand",
			},
		},
		{
			"$unwind": "$brand",
		},
		{
			"$match": bson.M{
				"brand._id": 4,
			},
		},
		{
			"$lookup": bson.M{
				"from":         categoryCollection.Name(),
				"localField":   "category_id",
				"foreignField": "_id",
				"as":           "category",
			},
		},
		{
			"$unwind": "$category",
		},
		{
			"$lookup": bson.M{
				"from":         variantsCollection.Name(),
				"localField":   "_id",
				"foreignField": "product_id",
				"as":           "variants",
			},
		},
	}
	// Execute the aggregation pipeline
	cursor, err := productCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// Process the results
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// Process the joined documents
		// You can access the fields using result["field_name"]

		fmt.Println(result)
	}
```
