(ns aoc.day-03
  (:require
   [clojure.string :as s]))

(defn bin->long [binary]
  (Long/parseLong binary  2))

(defn transpose [xs]
  (apply map vector xs))

(defn sort-digits
  "Returns [least-common most-common]"
  [digits]
  (->> digits
       (frequencies)
       (sort-by val)
       (map key)))


(defn part-1
  "Day 03 Part 1"
  [input]
  (->> input
       (s/split-lines)
       (map vec)
       (transpose)
       (map sort-digits)
       (transpose)
       (map s/join)
       (map bin->long)
       (apply *)))

(defn part-2
  "Day 03 Part 2"
  [input]
  input)
