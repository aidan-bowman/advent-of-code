(defun parse-file (file)
  (mapcar #'(lambda (s) (coerce s 'list)) (uiop:read-file-lines file)))

(defun check-bounds (map coords)
  (and (>= (first coords) 0)
       (>= (second coords) 0)
       (< (first coords) (length (first map)))
       (< (second coords) (length map))))

(defvar *antinodes* nil)

(defun add-antinode (map coords)
  (if (check-bounds map coords) (pushnew coords *antinodes* :test #'equal)))

(defun get-antinode (scoord ecoord)
  (list (+ (* (- (first ecoord) (first scoord)) 2) (first scoord))
        (+ (* (- (second ecoord) (second scoord)) 2) (second scoord))))

(defun get-map-coord (map coords)
  (nth (first coords) (nth (second coords) map)))

(defun antis-from-node (map coords)
  (let ((char (get-map-coord map coords)))
    (if (not (equal char #\.))
        (loop for y from 0 to (length map)
              do (loop for x from 0 to (length (first map))
                       if (not (equal (list x y) coords))
                       do (if (equal (get-map-coord map (list x y)) char) (add-antinode map (get-antinode coords (list x y)))))))))

(defun check-all-nodes (map func)
  (loop for y from 0 to (length map)
        do (loop for x from 0 to (length (first map))
                 do (funcall func map (list x y)))))

(defun part1 (file)
  (setf *antinodes* nil)
  (check-all-nodes (parse-file file) #'antis-from-node)
  (length *antinodes*))

(defun handle-antinodes (map scoord ecoord)
  (let ((dx (- (first ecoord) (first scoord)))
        (dy (- (second ecoord) (second scoord))))
    (loop for step from 0
          if (not (add-antinode map (list (- (first scoord) (* dx step)) (- (second scoord) (* dy step))))) do (return))
    (loop for step from 0
          if (not (add-antinode map (list (+ (first ecoord) (* dx step)) (+ (second ecoord) (* dy step))))) do (return))))

(defun antis-from-node-2 (map coords)
  (let ((char (get-map-coord map coords)))
    (if (not (equal char #\.))
        (loop for y from 0 to (length map)
              do (loop for x from 0 to (length (first map))
                       if (not (equal (list x y) coords))
                       do (if (equal (get-map-coord map (list x y)) char) (handle-antinodes map coords (list x y))))))))

(defun part2 (file)
  (setf *antinodes* nil)
  (check-all-nodes (parse-file file) #'antis-from-node-2)
  (length *antinodes*))
