(ns aoc.day-02
  (:require
   [clojure.string :as s]
   [aoc.utils :as u]))

; part 1
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

; part 2
(defn parse-command-2 [s]
  (let [[d x] (s/split s #" ")]
    [d (u/parse-int x)]))

; first argument is a collection of state [pos depth aim]
; second argument is a collection of commands [dir mag]
(defn advance-submarine [[pos depth aim] [dir mag]]
  (case dir
    "down" [pos depth (+ aim mag)]
    "up"   [pos depth (- aim mag)]
    "forward" [(+ pos mag) (+ depth (* aim mag)) aim]))

(defn part-2
  "Day 02 Part 2"
  [input]
  (let [[pos depth] (reduce advance-submarine
                            [0 0 0]
                            (->> input
                                 (s/split-lines)
                                 (map parse-command-2)))]
    (* pos depth)))
