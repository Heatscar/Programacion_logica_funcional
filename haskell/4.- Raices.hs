raices (a, b ,c) = if d < 0 then error "0" else (x, y)
    where
        x = e + sqrt d / (2 * a)
        y = e - sqrt d / (2 * a)
        d = b * b - 4 * a * c
        e = - b / (2 * a)




main = do putStr "Dame el primer valor: "

          a <- getLine
          putStr "Dame el segundo valor: "
          b <- getLine
          putStr "Dame el tercer valor : "
          c <- getLine

          putStr "La raiz es: "
          putStrLn (show (raices (read a, read b,read c)))
