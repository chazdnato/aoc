(ns aoc.day-02
  (:require
   [clojure.string :as s]
   [aoc.utils :as u]))

(def directions
  {"forward" [1 0]
   "down" [0 1]
   "up" [0 -1]})

(defn- parse-command [s]
  (let [[d x] (s/split s #" ")]
    (repeat (u/parse-int x)
            (get directions d))))

(defn part-1
  "Day 02 Part 1"
  [input]
  (let [pairs (->> input
                   (s/split-lines)
                   (mapcat parse-command))]
    (* (reduce + (map first pairs))
       (reduce + (map second pairs)))))

(defn part-2
  "Day 02 Part 2"
  [input]
  input)
