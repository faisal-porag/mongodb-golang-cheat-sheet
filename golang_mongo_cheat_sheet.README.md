


> document that have atleast one product with name apple and quantity 15 in json array using mongo in go 

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
