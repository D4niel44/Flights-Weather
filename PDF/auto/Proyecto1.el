(TeX-add-style-hook
 "Proyecto1"
 (lambda ()
   (TeX-add-to-alist 'LaTeX-provided-package-options
                     '(("babel" "spanish") ("inputenc" "utf8") ("fontenc" "T1") ("geometry" "top=3cm" "bottom=3cm" "left=2cm" "right=2cm" "heightrounded" "")))
   (TeX-run-style-hooks
    "latex2e"
    "report"
    "rep10"
    "babel"
    "inputenc"
    "fontenc"
    "amsthm"
    "subcaption"
    "graphicx"
    "amsmath"
    "geometry"))
 :latex)

