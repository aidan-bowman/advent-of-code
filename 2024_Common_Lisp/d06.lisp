(defun parse-file-into-lists (file)
  (mapcar #'(lambda (s) (coerce s 'list)) (uiop:read-file-lines file)))

;; we'll represent the guard as (x y dir) where dir is 0-3 starting north and going clockwise.

(setf *visitedlist* nil)

(defun find-pos-guard (map)
  (block loops
    (loop for x from 0 to (length (first map))
          do (loop for y from 0 to (length map)
                   if (equal (nth x (nth y map)) #\^) do (return-from loops (list x y 0))
                   if (equal (nth x (nth y map)) #\>) do (return-from loops (list x y 1))
                   if (equal (nth x (nth y map)) #\v) do (return-from loops (list x y 2))
                   if (equal (nth x (nth y map)) #\<) do (return-from loops (list x y 3))))))

(defun get-guard-view-coords (guard)
  (case (third guard)
    (3 (list (- (first guard) 1) (second guard) 3))
    (2 (list (first guard) (+ (second guard) 1) 2))
    (1 (list (+ (first guard) 1) (second guard) 1))
    (0 (list (first guard) (- (second guard) 1) 0))))

(defun val-at-coords (map coords)
  (nth (first coords) (nth (second coords) map)))

(defun visit-coord (coord)
  (setf *visitedlist* (adjoin (list (first coord) (second coord)) *visitedlist* :test #'equal)))

(defun out-of-bounds (map coords)
  (or (< (first coords) 0)
      (< (second coords) 0)
      (> (first coords) (- (length (first map)) 1))
      (> (second coords) (- (length map) 1))))

(defun run-step (map guard)
  (if (equal (val-at-coords map (get-guard-view-coords guard)) #\#)
      (if (equal (val-at-coords map (get-guard-view-coords (list (first guard) (second guard) (mod (+ 1 (third guard)) 4)))) #\#)
          (get-guard-view-coords (list (first guard) (second guard) (mod (+ (third guard) 2) 4)))
          (get-guard-view-coords (list (first guard) (second guard) (mod (+ (third guard) 1) 4))))
      (get-guard-view-coords guard)))

(defun run-sim (map)
  (let ((guard (find-pos-guard map)))
    (loop
      do (visit-coord guard)
      do (setf guard (run-step map guard))
      if (out-of-bounds map guard) do (return (length *visitedlist*)))))

(defun part1 (file)
  (setf *visitedlist* nil)
  (run-sim (parse-file-into-lists file)))

(defun visit-coord-loop (coord)
  (if (member coord *visitedlist* :test #'equal)
      nil
      (setf *visitedlist* (cons coord *visitedlist*))))

(defun add-wall (map coord)
  (let ((temp (copy-tree map)))
    (replace (nth (second coord) temp) "#" :start1 (first coord) :end1 (+ (first coord) 1))
    temp))

(defun run-sim-loop (map)
  (setf *visitedlist* nil)
  (let ((guard (find-pos-guard map)))
    (loop
      if (not (visit-coord-loop guard)) do (return t)
      if (out-of-bounds map (get-guard-view-coords guard)) do (return nil)
      do (setf guard (run-step map guard)))))

(defun run-tests (map)
  (run-sim map)
  (let* ((guardpos (find-pos-guard map))
        (increment 0)
        (listwalls (remove (list (first guardpos) (second guardpos)) (copy-tree *visitedlist*) :test #'equal)))
    (loop for val in listwalls
          if (run-sim-loop (add-wall map val)) do (setf increment (+ increment 1)))
    increment))

(defun part2 (file)
  (run-tests (parse-file-into-lists file)))

