;;; split input into 2 files - 1 for rules, 1 for updates

(defun parse-rules (file)
  (let ((strings (uiop:read-file-lines file)))
    (mapcar #'(lambda (a) (cons (parse-integer a :start 0 :end 2) (parse-integer a :start 3 :end 5))) strings)))

(defun parse-updates (file)
  (mapcar #'(lambda (a) (mapcar #'parse-integer (uiop:split-string a :separator ","))) (uiop:read-file-lines file)))

(defun verify-single-rule (rule update)
  (member (car rule) (member (cdr rule) update)))

(defun verify-all-rules (rules update)
  (notany #'(lambda (a) (verify-single-rule a update)) rules))

(defun part1 (rfile ufile)
  (let ((rules (parse-rules rfile)))
    (reduce #'+ (mapcar #'(lambda (n) (nth (/ (- (length n) 1) 2) n))
                        (remove-if-not #'(lambda (n) (verify-all-rules rules n)) (parse-updates ufile))))))

(defun check-order (rules a b)
  (find-if #'(lambda (n) (equal n (cons a b))) rules))

(defun sort-update (rules update)
  (sort update #'(lambda (a b) (check-order rules a b))))

(defun part2 (rfile ufile)
  (let ((rules (parse-rules rfile)))
    (reduce #'+ (mapcar #'(lambda (n) (nth (/ (- (length n) 1) 2) (sort-update rules n)))
                        (remove-if #'(lambda (n) (verify-all-rules rules n)) (parse-updates ufile))))))
