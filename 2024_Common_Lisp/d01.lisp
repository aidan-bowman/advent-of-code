(defun read-first-list ()
  (mapcar #'parse-integer (mapcar #'first (mapcar #'(lambda (n) (uiop:split-string n :separator " ")) (uiop:read-file-lines "/home/cheesepizza2/Downloads/input")))))

(defun read-second-list ()
  (mapcar #'parse-integer (mapcar #'fourth (mapcar #'(lambda (n) (uiop:split-string n :separator " ")) (uiop:read-file-lines "/home/cheesepizza2/Downloads/input")))))

(defun part1 ()
  (apply #'+ (mapcar #'abs (mapcar #'- (sort (read-first-list) #'<) (sort (read-second-list) #'<)))))

(defun part2 ()
  (let ((seclist (read-second-list)))
    (apply #'+ (mapcar #'(lambda (n) (* n (length (remove-if-not #'(lambda (k) (equal k n)) seclist)))) (read-first-list)))))
