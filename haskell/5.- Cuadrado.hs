cuadrado (n) = sum $ take n [1,3..]

help :: IO ()
help = do
	a <- getLine
	let i = 1
	if i == a
          then
          else
	      putStrLn (show (cuadrado (read i)))
	      let i += 1
              help

main :: IO ()
main = do putStr "Dame el valor de N: "
          help
