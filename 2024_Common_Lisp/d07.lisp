(defun parse-file (file)
  (mapcar #'(lambda (n) (mapcar #'(lambda (i) (parse-integer i :junk-allowed t)) (uiop:split-string n :separator " "))) (uiop:read-file-lines file)))

(defun check-equation (li ops ans)
  (let ((curval (first li)))
    (loop for x from 0 to (- (length li) 2)
          do (if (zerop (logand ops (ash 1 x)))
                 (setf curval (* curval (nth (+ 1 x) li)))
                 (setf curval (+ curval (nth (+ 1 x) li))))
          if (> curval ans) do (return nil))
    (equal curval ans)))

(defun check-values (li)
  (loop for ops from 0 to (- (expt 2 (- (length li) 1)) 1)
        if (check-equation (cdr li) ops (first li)) do (return t)))

(defun part1 (file)
  (reduce #'+ (mapcar #'first (remove-if-not #'check-values (parse-file file)))))

(defun check-equation-2 (li cur ans)
  (cond ((null li) (equal cur ans))
        ((> cur ans) nil)
        (t (or (check-equation-2 (cdr li) (+ cur (car li)) ans)
               (check-equation-2 (cdr li) (* cur (car li)) ans)
               (check-equation-2 (cdr li) (parse-integer (format nil "~D~D" cur (car li))) ans)))))

(defun check-values-2 (li)
  (check-equation-2 (cdr li) 0 (car li)))

(defun part2 (file)
  (reduce #'+ (mapcar #'first (remove-if-not #'check-values-2 (parse-file file)))))
