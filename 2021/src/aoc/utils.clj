(ns aoc.utils
  (:require [clojure.string :refer [split-lines]])
  (:require [clojure.java.io :refer [resource]]))


(defn parse-int
  "parses string to integer"
  [number-string]
  (Integer/parseInt number-string 10))

(defn lines
  "reads lines from a resource"
  [name]
  (->> name
       (resource)
       (slurp)
       (split-lines)))
