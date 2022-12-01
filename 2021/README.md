# Advent of Code Clojure

A batteries included starter pack for participating in [Advent of Code](https://www.adventofcode.com) using Clojure!

This was cloned from git@github.com:mhanberg/advent-of-code-clojure-starter.git

(This projects uses [lein](https://github.com/technomancy/leiningen)).

## Usage

There are 9 namespaces, 9 input files 9 example input files, 9 tests, and 18 `lein` tasks.

1. Fill in the tests with the example solutions.
1. Write your implementation.
1. Fill in the final problem inputs into the `lein` task and run `lein run d01.p1`!

```clojure
(ns advent-of-code.day-01)

(defn part-1
  "Day 01 Part 1"
  [input]
  input)

(defn part-2
  "Day 01 Part 2"
  [input]
  input)
```

```clojure
(ns advent-of-code.core-test
  (:require [clojure.test :refer [deftest testing is]]
            [advent-of-code.day-01 :refer [part-1 part-2]]
            [clojure.java.io :refer [resource]]))

(deftest part1
  (let [expected nil]
    (is (= expected (part-1 (slurp (resource "day-1-example.txt")))))))

(deftest part2
  (let [expected nil]
    (is (= expected (part-2 (slurp (resource "day-1-example.txt")))))))
```
