(defun parse-file (file)
  (mapcar #'digit-char-p (coerce (uiop:read-file-line file) 'list)))

(defun file-checksum (index size id)
  (* id
     (+ (* index size) (/ (* size (- size 1)) 2))))

(defun rm-first-list (list)
  (setf list (cddr list)))

(defun get-fid-front (lsize)
  (/ (- lsize 1) 2))

(defun rpos (pos list)
  (- (length list) pos 1))

(defun rnth (pos list)
  (nth (rpos pos list) list))

(defun step-disk (index pos list)
  (if (equal (+ pos 1) (length list))
      (file-checksum index (first list) (/ pos 2))
      (+ (file-checksum index (rnth pos list) (/ pos 2))
         (step-space (+ index (rnth pos list)) (+ pos 1) list))))

(defun step-space (index pos list)
  (if (equal (+ pos 2) (length list))
      (step-disk index (+ pos 1) list)
      (cond ((> (first list) (rnth pos list)) (progn (setf (first list) (- (first list) (rnth pos list)))
                                                     (+ (file-checksum index (rnth pos list) (get-fid-front (length list)))
                                                        (step-disk (+ index (rnth pos list)) (+ pos 1) list))))
            ((equal (first list) (rnth pos list)) (+ (file-checksum index (first list) (get-fid-front (length list)))
                                                     (step-disk (+ index (rnth pos list)) (+ pos 1) (cddr list))))
            (t (let ((temp (file-checksum index (first list) (get-fid-front (length list))))
                     (size-handled (first list)))
                 (setf (nth (rpos pos list) list) (- (rnth pos list) size-handled))
                 (+ temp
                    (step-space (+ index size-handled) pos (cddr list))))))))

(defun part1 (file)
  (let ((map (reverse (parse-file file))))
    (step-disk 0 0 map)))

(defun split-map (map)
  (let ((index 0)
        (id 0))
    (loop for (a b) on map by #'cddr
          collect (list a index id) into filelist
          do (setf index (+ index a) id (+ id 1))
          if (not (null b))
          collect (list b index) into spacelist and
          do (setf index (+ index b))
          finally (return (list filelist spacelist)))))

(defun file-into-space (file spacelist)
  (if (or (null spacelist)
          (> (second (first spacelist)) (second file)))
      (file-checksum (second file) (first file) (third file))
      (if (< (first (first spacelist)) (first file))
             (file-into-space file (cdr spacelist))
             (let ((result (file-checksum (second (first spacelist)) (first file) (third file))))
               (setf (first (first spacelist)) (- (first (first spacelist)) (first file)))
               (setf (second (first spacelist)) (+ (second (first spacelist)) (first file)))
               result))))

(defun empty-space (spacelist)
  (if (null spacelist)
      nil
      (if (zerop (first (first spacelist)))
             (empty-space (cdr spacelist))
             spacelist)))

(defun rec-handle-files (files spacelist)
  (if (null files)
      0
      (+ (file-into-space (first files) spacelist)
         (rec-handle-files (cdr files) (empty-space spacelist)))))

(defun part2temp (map)
  (let ((smap (split-map map)))
    (rec-handle-files (reverse (first smap)) (second smap))))

(defun part2 (file)
  (let ((smap (split-map (parse-file file))))
    (rec-handle-files (reverse (first smap)) (second smap))))
