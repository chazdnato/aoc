(ns aoc.day-01-test
  (:require [clojure.test :refer [deftest testing is]]
            [aoc.day-01 :refer [part-1 part-2]]
            [aoc.core :refer [lines parse-int]]))

(deftest part1
  (testing "day 01 part 1"
    (let [expected 7]
      (is (= expected (part-1 (map parse-int (lines "day-01-example.txt"))))))))

(deftest part2
  (testing "day 01 part 2"
    (let [expected 5]
      (is (= expected (part-2 (map parse-int (lines "day-01-example.txt"))))))))
