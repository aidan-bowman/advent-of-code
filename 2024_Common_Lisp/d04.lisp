(ql:quickload "cl-ppcre")

(defun get-file-lists (file)
  (mapcar #'(lambda (s) (coerce s 'list)) (uiop:read-file-lines file)))

(defun vert-list (ls)
  (loop for x from 0 to (- (length (first ls)) 1)
        collect (mapcar #'(lambda (n) (nth x n)) ls)))

(defun diag-l-list (ls)
  (loop for x from 0 to (- (+ (length ls) (length (first ls))) 2)
        collect (loop for y from (min x (- (length (first ls)) 1)) downto (max 0 (+ (- x (length ls)) 1))
                      collect (nth y (nth (- x y) ls)))))

(defun diag-r-list (ls)
  (loop for x from 0 to (- (+ (length ls) (length (first ls))) 2)
        collect (loop for y from (max 0 (+ (- x (length ls)) 1)) to (min x (- (length (first ls)) 1))
                      collect (nth y (nth (- (length ls) (- x y) 1) ls)))))

(defun check-string-xmas-count (s)
  (/ (+ (length (ppcre:all-matches "XMAS" s)) (length (ppcre:all-matches "SAMX" s))) 2))

(defun check-lists-for-xmas (ls)
  (apply #'+ (mapcar #'(lambda (l) (check-string-xmas-count (coerce l 'string))) ls)))

(defun part1 (file)
  (let ((ls (get-file-lists file)))
    (+ (check-lists-for-xmas ls)
       (check-lists-for-xmas (vert-list ls))
       (check-lists-for-xmas (diag-l-list ls))
       (check-lists-for-xmas (diag-r-list ls)))))

(defun check-pos-for-x-mas (ls x y)
  (and (equal (nth x (nth y ls)) #\A)
       (or (and (equal (nth (- x 1) (nth (- y 1) ls)) #\M)
                (equal (nth (+ x 1) (nth (+ y 1) ls)) #\S))
           (and (equal (nth (- x 1) (nth (- y 1) ls)) #\S)
                (equal (nth (+ x 1) (nth (+ y 1) ls)) #\M)))
       (or (and (equal (nth (- x 1) (nth (+ y 1) ls)) #\M)
                (equal (nth (+ x 1) (nth (- y 1) ls)) #\S))
           (and (equal (nth (- x 1) (nth (+ y 1) ls)) #\S)
                (equal (nth (+ x 1) (nth (- y 1) ls)) #\M)))))

(defun part2 (file)
  (let ((ls (get-file-lists file)))
    (loop for x from 1 to (- (length (first ls)) 2)
          summing (loop for y from 1 to (- (length ls) 2)
                        count (check-pos-for-x-mas ls x y) into total
                        finally (return total)) into total
          finally (return total))))
