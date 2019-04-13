suma (a,b) (c,d) = (a + c , b + d)

sumas = do putStr "Dame la parte real 1: "
           a <- getLine
           putStr "Dame la parte imaginaria 1: "
           b <- getLine
           putStr "Dame la parte real 2: "
           c <- getLine
           putStr "Dame la parte imaginaria 2: "
           d <- getLine
           putStr "La suma es: "
           putStrLn (show (suma (read a, read b) (read c, read d)))
