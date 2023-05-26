# mongodb-golang-cheat-sheet
mongodb &amp; golang

---

Mongodb-Aggregation
----------------------------------------------------------------------------------------------------------------------------


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





