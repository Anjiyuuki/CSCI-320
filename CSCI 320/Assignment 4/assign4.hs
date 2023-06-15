--1 Piecewise Functions and Recursion - A, B
piece' :: Double -> Double 
--will take in a double value and return a double value
piece' x -- the input value
  | x < -1 = 2 * x + 1
  | x >= -1 && x <= 3 = -2
  | x > 3 = -3 * x + 7

--two argument type a and return a (float or double)
exp' :: (Floating z, Ord z) => z -> z -> z
exp' x y
  | y == 0    = 1 --base case
  | y < 0     = 1 / exp' x (-y) -- if y neg, 1 / (x^(-y))
  | otherwise = x * exp' x (y-1) -- if y pos, x * (x^(y-1))


--2 Functional Composition
f xs = map (^2) xs
g xs = sum xs
gf = g . f
x xs = gf xs

comp' = do
  input <- getLine
  let xs = map read (words input)
  print (x xs)

--3 List Operations - A, B, C
count :: (Int -> Bool) -> [Int] -> Int
count criteria list = sum [1 | element <- list, criteria element == True]

count'filter :: (Int -> Bool) -> [Int] -> Int
count'filter criteria list = length (filter criteria list)

count'lc :: (Int -> Bool) -> [Int] -> Int
count'lc criteria list = length [element | element <- list, criteria element == True]

--4 Currying
closure :: (Int -> Bool) -> [Int] -> Int
closure evaluator = count evaluator

-- return true or false if the value is greater than 10
isValueGreaterThan10 :: Int -> Bool
isValueGreaterThan10 value = value > 10

--curried function
countValuesGreaterThan10 :: [Int] -> Int
countValuesGreaterThan10 = closure isValueGreaterThan10

--5 Extra Credit A, B
powers' :: Int -> [Int]
-- infinite list of pos int 
powers' base = [base^i | i <- [1..]] 

count'' :: (Int -> Bool) -> [Int] -> Int
--length determines the number of elements in the list
-- eval will return true or false
count'' eval list = length filteredList
-- will only return if it eval to true
  where filteredList = filter eval list
