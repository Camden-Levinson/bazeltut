
from random import randint

from project.calculator.calculator import Calculator

my_calculator = Calculator()

def main():
    num1 = randint(0, 100)
    num2 = randint(0, 100)
    message = "{} + {} = {}".format(num1, num2, my_calculator.add(num1, num2))
    return message

if __name__ == '__main__':
    print(main())