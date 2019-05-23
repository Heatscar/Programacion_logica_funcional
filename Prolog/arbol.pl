hombre(honorio).
hombre(agustin).
hombre(miguel).
hombre(antonio).
hombre(pedro).
hombre(george).
hombre(fred).
hombre(alonso).
hombre(alejandro).
hombre(carlos).
hombre(hector).
hombre(hectorh).

mujer(raquel).
mujer(marcelina).
mujer(alejandra).
mujer(leonor).
mujer(andrea).
mujer(lucy).
mujer(molly).
mujer(angelina).
mujer(roxanne).
mujer(beatriz).
mujer(reyna).
mujer(maria).
mujer(erika).

progenitor(raquel,agustin).
progenitor(raquel,pedro).
progenitor(raquel,george).
progenitor(raquel,beatriz).
progenitor(raquel,hector).
progenitor(honorio,agustin).
progenitor(honorio,pedro).
progenitor(honorio,george).
progenitor(honorio,beatriz).
progenitor(honorio,hector).
progenitor(agustin,alejandra).
progenitor(agustin,leonor).
progenitor(agustin,miguel).
progenitor(marcelina,alejandra).
progenitor(marcelina,leonor).
progenitor(marcelina,miguel).
progenitor(pedro,lucy).
progenitor(pedro,molly).
progenitor(andrea,lucy).
progenitor(andrea,molly).
progenitor(george,fred).
progenitor(george,roxanne).
progenitor(angelina,fred).
progenitor(angelina,roxanne).
progenitor(beatriz,alejandro).
progenitor(beatriz,reyna).
progenitor(beatriz,carlos).
progenitor(alonso,alejandro).
progenitor(alonso,reyna).
progenitor(alonso,carlos).
progenitor(hector,hectorh).
progenitor(hector,erika).
progenitor(maria,hectorh).
progenitor(maria,erika).

pareja(honorio,molly).
pareja(molly,honorio).
pareja(agustin,marcelina).
pareja(marcelina,agustin).
pareja(pedro,andrea).
pareja(andrea,pedro).
pareja(george,angelina).
pareja(angelina,george).
pareja(beatriz,alonso).
pareja(alonso,beatriz).
pareja(hector,maria).
pareja(maria,hector).
pareja(alejandra,antonio).
pareja(antonio,alejandra).

padre(X,Y):-hombre(X),progenitor(X,Y).
madre(X,Y):-mujer(X),progenitor(X,Y).
hermanos(X,Y):-progenitor(Z,X),progenitor(Z,Y), not(X==Y).
hermano(X,Y):-hombre(X),hermanos(X,Y).
hermana(X,Y):-mujer(X),hermanos(X,Y).
esposo(X,Y):-pareja(X,Y),hombre(X).
esposa(X,Y):-pareja(X,Y),mujer(X).
suegro(X,Y):-padre(X,Z),pareja(Y,Z).
suegra(X,Y):-madre(X,Z),esposos(Y,Z).
yerno(X,Y):-suegro(Y,X);suegra(Y,X),hombre(X).
nuera(X,Y):-suegro(Y,X);suegra(Y,X),mujer(X).
cuñados(X,Y):-((pareja(X,Z),hermanos(Z,Y));(pareja(Y,Z),hermanos(Z,X))).
cuñado(X,Y):-cuñados(X,Y),hombre(X).
cuñada(X,Y):-cuñados(X,Y),mujer(X).
abuelo(X,Y):-progenitor(Z,Y),padre(X,Z).
abuela(X,Y):-progenitor(Z,Y),madre(X,Z).
nieto(X,Y):-progenitor(Y,Z),progenitor(Z,X),hombre(X).
nieta(X,Y):-progenitor(Y,Z),progenitor(Z,X),mujer(X).
tio(X,Y):-progenitor(Z,Y),(hermano(X,Z);cuñado(X,Z)).
tia(X,Y):-progenitor(Z,Y),(hermana(X,Z);cuñada(X,Z)).
primo(X,Y):-progenitor(Z,X),progenitor(W,Y),hermanos(Z,W),hombre(X).
prima(X,Y):-progenitor(Z,X),progenitor(W,Y),hermanos(Z,W),mujer(X).


% Pruebas
% primo(miguel,erika).
% nieto(carlos,honorio).
% tio(agustin,lucy).
% hermanos(carlos,alejandro).
% padre(hector,hectorh).
% madre(maria,hectorh).
% esposo(hector,maria).
% abuelo(honorio,alejandro).