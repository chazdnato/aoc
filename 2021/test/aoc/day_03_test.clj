(ns aoc.day-03-test
  (:require [clojure.test :refer [deftest testing is]]
            [aoc.day-03 :refer [part-1 part-2]]
            [clojure.java.io :refer [resource]]))

(deftest part1
  (testing "day 03 part 1"
    (let [expected 198]
      (is (= expected (part-1 (slurp (resource "day-03-example.txt"))))))))

(deftest part2
  (testing "day 03 part 2"
    (let [expected 230]
      (is (= expected (part-2 (slurp (resource "day-03-example.txt"))))))))
