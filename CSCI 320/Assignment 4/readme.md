Angie Bui
2/9/23

# Assignment 4 - Haskell

Haskell function
Dealing with Piecewise Functions and Recursion, Functional Composition,
List Operations, Currying

## Running
Open folder in Visual Studios

From the command line:

```
ghci
:l assign4.hs

--- #1 a and b sample
piece' 10 
:r
exp' 2 3
:r

--- #2 sample
---after you type comp' and enter... type the list of number you want with a space then enter, don't need [])---
comp'
1 2 3 4 5
:r

--- #3 Sample A, B, C
isNegative x = x < 0
myList = [-3, -2, -1, 0, 1]

--- count function
count isNegative myList 
--- count'filter function
count'filter isNegative myList 
--- count'lc function
count'lc isNegative myList 


--- #4 Sample
:r
countValuesGreaterThan10 [1, 2, 3, 5, 10, 11, 14, 20]

--- #5 Sample A, B
:r 
take 4 (powers' 2)

:r
let isValueGreaterThan10 value = value >10
let myL = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]
count'' isValueGreaterThan10 myL
```

