# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

from ctypes import *
import json
import sys

# Load Shared Object
# gotmpl.so must be in the LD_LIBRARY_PATH
# By default, LD_LIBRARY_PATH does not include current directory.
# You can add current directory with LD_LIBRARY_PATH=. python test.py
lib = cdll.LoadLibrary("gotmpl.so")

# Define Function Definitions
version = lib.Version
version.argtypes = []
version.restype = c_char_p

render = lib.Render
render.argtypes = [c_char_p, c_char_p, c_char_p, POINTER(c_char_p)]
render.restype = c_char_p

# Define input and output variables
# Output must be a ctypec_char_p
tmpl = "Sum: {{ .values | sum }}\nMax: {{ .values | max }}\nValues: {{ .values | join \", \" }}"
ctx = {
  "values": [1, 2, 3, 4]
}
output_string_pointer = c_char_p()

print version()

print "# Template"
print tmpl

err = render(tmpl, json.dumps(ctx), "json", byref(output_string_pointer))
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

# Convert from ctype to python string
output_string = output_string_pointer.value

# Print output to stdout
print "# Rendered"
print output_string
