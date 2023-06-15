quickSort :: (Ord a) => [a] -> [a]
quickSort [] = []
quickSort (central:list) =
    let lowerSorted = quickSort [x | x <- list, x <= central]
        higherSorted = quickSort [x | x <- list, x > central]
    in lowerSorted ++ [central] ++ higherSorted
