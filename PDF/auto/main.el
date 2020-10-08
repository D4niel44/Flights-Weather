(TeX-add-style-hook
 "main"
 (lambda ()
   (TeX-add-to-alist 'LaTeX-provided-package-options
                     '(("babel" "spanish") ("inputenc" "utf8") ("fontenc" "T1") ("datetime" "nodayoftheweek") ("geometry" "top=3cm" "bottom=3cm" "left=2cm" "right=2cm" "heightrounded" "")))
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
    "datetime"
    "geometry")
   (TeX-add-symbols
    "mydate"))
 :latex)

