(ns aoc.day-02-test
  (:require [clojure.test :refer [deftest testing is]]
            [aoc.day-02 :refer [part-1 part-2]]
            [clojure.java.io :refer [resource]]))

(deftest part1
  (testing "day 02 part 1"
    (let [expected 150]
      (is (= expected (part-1 (slurp (resource "day-02-example.txt"))))))))

(deftest part2
  (testing "day 02 part 2"
    (let [expected 900]
      (is (= expected (part-2 (slurp (resource "day-02-example.txt"))))))))
