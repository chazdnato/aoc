(ns aoc.core
  (:require [aoc.day-01] [aoc.day-02] [aoc.day-03] [aoc.day-04] [aoc.day-05] [aoc.day-06] [aoc.day-07] [aoc.day-08] [aoc.day-09])
  (:require [clojure.java.io :refer [resource]])
  (:require [clojure.string :refer [split-lines]]))

(defn read-input
  [day]
  (slurp (resource day)))

(defn lines
  "reads lines from a resource"
  [name]
  (->> name
       (resource)
       (slurp)
       (split-lines)))

(defn parse-int
  "parses string to integer"
  [number-string]
  (Integer/parseInt number-string 10))


(defn -main
  "Used to dispatch tasks from the command line.
  
  lein run d01.p1"
  [part]
  (case part
    "d01.p1" (println (aoc.day-01/part-1 (map parse-int (lines "day-01.txt"))))
    "d01.p2" (println (aoc.day-01/part-2 (map parse-int (lines "day-01.txt"))))
    "d02.p1" (println (aoc.day-02/part-1 (read-input "day-02.txt")))
    "d02.p2" (println (aoc.day-02/part-2 (read-input "day-02.txt")))
    "d03.p1" (println (aoc.day-03/part-1 (read-input "day-03.txt")))
    "d03.p2" (println (aoc.day-03/part-2 (read-input "day-03.txt")))
    "d04.p1" (println (aoc.day-04/part-1 (read-input "day-04.txt")))
    "d04.p2" (println (aoc.day-04/part-2 (read-input "day-04.txt")))
    "d05.p1" (println (aoc.day-05/part-1 (read-input "day-05.txt")))
    "d05.p2" (println (aoc.day-05/part-2 (read-input "day-05.txt")))
    "d06.p1" (println (aoc.day-06/part-1 (read-input "day-06.txt")))
    "d06.p2" (println (aoc.day-06/part-2 (read-input "day-06.txt")))
    "d07.p1" (println (aoc.day-07/part-1 (read-input "day-07.txt")))
    "d07.p2" (println (aoc.day-07/part-2 (read-input "day-07.txt")))
    "d08.p1" (println (aoc.day-08/part-1 (read-input "day-08.txt")))
    "d08.p2" (println (aoc.day-08/part-2 (read-input "day-08.txt")))
    "d09.p1" (println (aoc.day-09/part-1 (read-input "day-09.txt")))
    "d09.p2" (println (aoc.day-09/part-2 (read-input "day-09.txt")))
    (println "not found")))

