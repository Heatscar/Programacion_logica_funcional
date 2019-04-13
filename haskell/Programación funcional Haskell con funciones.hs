--Distancia entre dos puntos
-- Ejemplo:  (-4, 3) (3, 2)
distancia:: (Float,Float) -> (Float,Float) -> Float
distancia (x1,y1) (x2,y2) = sqrt((x1-x2)^2+(y1-y2)^2)	


--Suma de dos numeros complejos
--Ejemplo: (3, -2) (6, -4)
sumaComplejos:: (Int,Int) ->(Int,Int) -> (Int,Int)
sumaComplejos (x1,x2) (y1,y2) = (x1+y1,x2+y2)


--Producto de dos numeros complejos
productoComplejos:: (Int,Int) -> (Int,Int) -> (Int,Int)
productoComplejos (x1,x2) (y1,y2) = (x1*y1-x2*y2,x1*y2+x2*y1) 


--Raice de las ecuaciones cuadraticas
--Ejemplo: raices 1 3 2
raices:: Float -> Float -> Float -> [Float]
raices a b c = [(-b+d)/t,(-b-d)/t]
	where 
		d = sqrt (b^2 - 4*a*c)
		t = 2*a