(ns aoc.day-01-test
  (:require [clojure.test :refer [deftest testing is]]
            [aoc.day-01 :refer [part-1 part-2]]
            [aoc.utils :as u]))

(deftest part1
  (testing "day 01 part 1"
    (let [expected 7]
      (is (= expected (part-1 (map u/parse-int (u/lines "day-01-example.txt"))))))))

(deftest part2
  (testing "day 01 part 2"
    (let [expected 5]
      (is (= expected (part-2 (map u/parse-int (u/lines "day-01-example.txt"))))))))
