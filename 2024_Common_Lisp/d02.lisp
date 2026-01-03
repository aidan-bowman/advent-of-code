(defun all-increasing (l)
  (every #'< l (append (cdr l) '(1000))))

(defun all-decreasing (l)
  (every #'> l (append (cdr l) '(0))))

(defun safe-diffs (l)
  (every #'(lambda (a b) (< (abs (- a b)) 4)) l (append (cdr l) (last l))))

(defun safe-report (l)
  (and (or (all-increasing l) (all-decreasing l)) (safe-diffs l)))

(defun parse-string (s)
  (mapcar #'parse-integer (uiop:split-string s :separator " ")))

(defun sum-reports (ls)
  (length (remove-if-not #'safe-report (mapcar #'parse-string ls))))

(defun rec-remove-nth (l n)
  (cond ((null l) nil)
        ((zerop n) (cdr l))
        (t (cons (car l) (rec-remove-nth (cdr l) (- n 1))))))

(defun dampen-check (l)
  (or
   (safe-report l)
   (loop for x from 0 to (- (length l) 1)
         when (safe-report (rec-remove-nth l x)) return t)))

(defun sum-reports-dampened (ls)
  (length (remove-if-not #'dampen-check (mapcar #'parse-string ls))))
