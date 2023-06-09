# MongoDB GoLang Cheat Sheet
### MongoDB &amp; GoLang

---

#### MongoDB --->>> Aggregation & Others Operations
----------------------------------------------------

##### CRUD OPERATIONS IN MONGO DB

> CREATE HAS TWO METHODS

```mongo
db.collectionName.insertOne(data, operations)
db.collectionName.insertMany(data, operations)
```

> READ HAS TWO METHODS

```mongo
db.collectionName.find(filter, operations)
db.collectionName.findOne(filter, operations)
```

> UPDATE HAS THREE METHODS


```mongo
db.collectionName.updateOne(filter, data, operations)
db.collectionName.updateMany(filter, data, operations)
db.collectionName.replaceOne(filter, data, operations)
```

> DELETE HAS TWO METHODS

```mongo
db.collectionName.deleteOne(filter, operations)
db.collectionName.deleteMany(filter, operations)
```


1.  Group by age and show the names in this group 
```mongo
db.teachers.aggregate(
	[
		{
			$group: {
_id: “$age”,
 names: {$push: “$name”} // Note: it is an array because of the list of names.
}
}
]
)
```

2.  Group by age and show the all information in this group 
```mongo
db.teachers.aggregate(
	[
		{
			$group: {
_id: “$age”,
 all_info: {$push: “$$ROOT”} 
}
}
]
)
```

3.  Give a count per age of male teacher 
```mongo
db.teachers.aggregate(
	[
		{
			$match: {
gender: “male”
}
},
		{
			$group: {
_id: “$age”,
 countOfTeacherInThisAgeGroup: {$sum: 1} 
}
}
]
)
```

4.  Give a count per age of male teachers and sort them by count in desc order
```mongo
db.teachers.aggregate(
	[
		{
			$match: {
gender: “male”
}
},
		{
			$group: {
_id: “$age”,
 countOfTeacherInThisAgeGroup: {$sum: 1} 
}
},
{
			$sort: {
countOfTeacherInThisAgeGroup: -1
}
}

]
)
```

5.  List All hobbies
```mongo
db.teachers.aggregate(
	[
		{
			$unwind: “$hobbies”
},
		{
			$group: {
_id: null,
 all_hobbies: {$push: “$hobbies”} 
}
},
]
)
```

6.  List All hobbies but unique ($addToSet -> for remove duplicates)
```mongo
db.teachers.aggregate(
	[
		{
			$unwind: “$hobbies”
},
		{
			$group: {
_id: null,
 all_hobbies: {$addToSet: “$hobbies”} 
}
},
])
```

7. Find average of scores of students whose age is greater than 20 
```mongo
db.students.aggregate([
{
$group: {
_id: null,
avgScore: {
$avg: {
$filter: {
input: "$scores”,
as: "score",
cond: { $gt: [ "Sage", 20] }
}
}
	}

}
}
])
```

$unwind - Example & Resource

[Read $unwind](https://www.bmc.com/blogs/mongodb-unwind/)


---


[Query in Array](https://www.youtube.com/watch?v=wPL2i8KHKTI&list=PLA3GkZPtsafZydhN4nP0h7hw7PQuLsBv1&index=23)



















