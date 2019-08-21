// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "gotmpl.h"

int
main(int argc, char **argv) {
    char *err;

    char *input_string = "Sum: {{ .values | sum }}\nMax: {{ .values | max }}\nValues: {{ .values | join \", \" }}";
    char *output_string;

    printf("# Template\n%s\n", input_string);

    err = Render(input_string, "{\"values\":[1,2,3,4]}", "json", &output_string);

    if (err != NULL) {
        fprintf(stderr, "error: %s\n", err);
        free(err);
        return 1;
    }

    printf("# Rendered:\n%s\n", output_string);

    return 0;
}
