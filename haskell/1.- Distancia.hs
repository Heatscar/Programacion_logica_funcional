distancia (x1 , y1) (x2 , y2) = sqrt (x'*x' + y'*y')
    where
      x' = x2 - x1
      y' = y2 - y1


main = do putStr "Dame la x 1: "
          a <- getLine
          putStr "Dame la y 1: "
          b <- getLine
          putStr "Dame la x 2: "
          c <- getLine
          putStr "Dame la y 2: "
          d <- getLine
          putStr "La distacia entre puntos es: "
          putStrLn (show (distancia (read a, read b) (read c, read d)))
