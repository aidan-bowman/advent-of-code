(ql:quickload "cl-ppcre")

(defun extract-muls (s)
  (ppcre:all-matches-as-strings "mul\\((\\d)+,(\\d)+\\)" s))

(defun calc-mul (s)
  (apply #'* (mapcar #'parse-integer (ppcre:all-matches-as-strings "(\\d)+" s))))

(defun part1 (input-file)
  (apply #'+ (mapcar #'calc-mul (extract-muls (uiop:read-file-string input-file)))))

(defun extract-insts (s)
  (ppcre:all-matches-as-strings "mul\\((\\d)+,(\\d)+\\)|do\\(\\)|don't\\(\\)" s))

(defun run-insts-do (ls)
  (cond ((null ls) 0)
        ((equal (first ls) "do()") (run-insts-do (rest ls)))
        ((equal (first ls) "don't()") (run-insts-dont (rest ls)))
        (t (+ (calc-mul (first ls)) (run-insts-do (rest ls))))))

(defun run-insts-dont (ls)
  (cond ((null ls) 0)
        ((equal (first ls) "do()") (run-insts-do (rest ls)))
        (t (run-insts-dont (rest ls)))))

(defun part2 (input-file)
  (run-insts-do (extract-insts (uiop:read-file-string input-file))))
