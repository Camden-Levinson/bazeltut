import unittest
from project.calculator.calculator import Calculator

class TestSum(unittest.TestCase):
    def test_sum(self):
        calculator = Calculator()
        self.assertEqual(calculator.add(1, 2), 3)
        self.assertEqual(calculator.subtract(5, 8), -3)
        self.assertEqual(calculator.multiple(4, 9), 36)
        self.assertEqual(calculator.divide(80, 5), 16)

if __name__ == '__main__':
    unittest.main()