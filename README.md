# testpro
testpro

задание 1 
один робот догоняет другого после того, как пройдет черный квадрат 
Answer

1.ML

2.IF FLAG GOTO 4

3.GOTO 1

4.ML

5.ML

6.GOTO 4

задание 2
нужно взять из вазс, на каторой написано "черные и белые". 1. Дастали черный шарик. => что эта ваза в черными шариками, у нас остались две вазы с надписями "черные" и "белые". тк нам известно, что они ложны, можно сделать вывод, что в вазе с надписью белые лежат "черные и белые", а с надписью "черыне" - белые. 
 2. Дастали белый шарик. => что эта ваза в белыми шариками, у нас остались две вазы с надписями "черные" и "белые". тк нам известно, что они ложны, можно сделать вывод, что в вазе с надписью "черные" лежат "черные и белые", а с надписью "белые" - черные.

3.запуск
git clone https://github.com/DiSShoRt/testpro

cd ./testpro

docker build -t go-app .

docker run  -p 5000:5000 -t go-app

