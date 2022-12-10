import turtle
import unittest
from time import sleep
from typing import Set, Tuple

direction_to_unit_vector = {
    "U": (0, 1),
    "D": (0, -1),
    "L": (-1, 0),
    "R": (1, 0),
}


def move_head(knot: Tuple[int, int], direction: str, magnitude: int) -> Tuple[int, int]:
    vector = direction_to_unit_vector[direction]
    return (knot[0] + vector[0] * magnitude, knot[1] + vector[1] * magnitude)


def out_of_range(x_2: int, x_1: int) -> bool:
    return abs(x_2 - x_1) > 1


def follow_tail(
    head_knot: Tuple[int, int],
    follow_knot: Tuple[int, int],
    move_set: Set[Tuple[int, int]],
):
    while out_of_range(head_knot[0], follow_knot[0]) or out_of_range(
        head_knot[1], follow_knot[1]
    ):
        x_del, y_del = 0, 0
        if head_knot[1] != follow_knot[1]:
            y_del = 1 if head_knot[1] > follow_knot[1] else -1
        if head_knot[0] != follow_knot[0]:
            x_del = 1 if head_knot[0] > follow_knot[0] else -1
        follow_knot = (follow_knot[0] + x_del, follow_knot[1] + y_del)
        if move_set is not None:
            print(follow_knot)
            # turtle.goto(follow_knot)
            move_set.add(follow_knot)

    return follow_knot


def get_tail_moves(input_str: str) -> int:
    knots = [(0, 0) for _ in range(10)]
    move_set = set()
    move_set.add((0, 0))
    # turtle.penup()
    # turtle.goto((0, 0))
    # turtle.pendown()
    for line in input_str.split("\n"):
        print(line)
        direction, magnitude = line.split(" ")
        for i in range(int(magnitude)):
            knots[0] = move_head(knots[0], direction, 1)
            for i in range(1, len(knots) - 1):
                knots[i] = follow_tail(knots[i - 1], knots[i], None)
            knots[9] = follow_tail(knots[8], knots[9], move_set)
        print(knots)
    # sleep(10)
    return len(move_set)


class PartTwo(unittest.TestCase):
    def test_part_two_small_sample(self):
        directions = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""
        self.assertEqual(get_tail_moves(directions), 1)

    def test_part_two_bigger_sample(self):
        directions = """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"""
        self.assertEqual(get_tail_moves(directions), 36)

    def test_part_two_final(self):
        with open("day9_input.txt") as f:
            directions = f.read()
            tail_moves = get_tail_moves(directions)
            self.assertEqual(5, tail_moves)


if __name__ == "__main__":
    unittest.main()
