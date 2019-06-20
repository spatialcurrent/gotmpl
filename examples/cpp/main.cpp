// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <iostream>
#include <string>
#include <cstring>
#include "gotmpl.h"

// render is a example of a C++ function that can render templates using some std::string variables.
// In production, you would want to write the function definition to match the use case.
char* render(std::string tmpl, std::string ctx, std::string format, char** output_string_c) {

  char *tmpl_c = new char[tmpl.length() + 1];
  std::strcpy(tmpl_c, tmpl.c_str());

  char *ctx_c = new char[ctx.length() + 1];
  std::strcpy(ctx_c, ctx.c_str());

  char *format_c = new char[format.length() + 1];
  std::strcpy(format_c, format.c_str());

  char *err = Render(tmpl_c, ctx_c, format_c, output_string_c);

  free(tmpl_c);
  free(ctx_c);
  free(format_c);

  return err;

}

int main(int argc, char **argv) {

  // Since Go requires non-const values, we must define our parameters as variables
  // https://stackoverflow.com/questions/4044255/passing-a-string-literal-to-a-function-that-takes-a-stdstring
  std::string tmpl("Sum: {{ .values | sum }}\nMax: {{ .values | max }}\nValues: {{ .values | join \", \" }}");
  std::string ctx("{\"values\":[1,2,3,4]}");
  std::string format("json");
  char *output_char_ptr;

  // Write version to stderr
  //std::cout << "Version" << Version() <<std::endl;

  // Write input to stderr
  std::cout << "Template:" << std::endl << tmpl << std::endl;

  char *err = render(tmpl, ctx, format, &output_char_ptr);
  if (err != NULL) {
    // Write output to stderr
    std::cerr << std::string(err) << std::endl;
    // Return exit code indicating error
    return 1;
  }
  std::string output_string = std::string(output_char_ptr);

  // Write output to stdout
  std::cout << "Rendered:" << std::endl << output_string << std::endl;

  // Return exit code indicating success
  return 0;
}
