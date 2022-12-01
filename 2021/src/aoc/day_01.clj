(ns aoc.day-01)

(defn part-1
  "Day 01 Part 1"
  [input]

  (->> input
       (partition 2 1)
       (filter #(apply < %))
       (count)))

(defn part-2
  "Day 01 Part 2"
  [input]
  (->> input
       (partition 3 1)
       (map #(apply + %))
       (partition 2 1)
       (filter #(apply < %))
       (count)))
