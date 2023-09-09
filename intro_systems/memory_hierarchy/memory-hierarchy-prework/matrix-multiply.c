/*
Naive code for multiplying two matrices together.

There must be a better way!
*/

#include <stdio.h>
#include <stdlib.h>

/*
  A naive implementation of matrix multiplication.

  DO NOT MODIFY THIS FUNCTION, the tests assume it works correctly, which it
  currently does
*/
void matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols,
                     int b_cols) {
  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      C[i][j] = 0;
      for (int k = 0; k < a_cols; k++) C[i][j] += A[i][k] * B[k][j];
    }
  }
}

void fast_matrix_multiply(double **c, double **a, double **b, int a_rows,
                          int a_cols, int b_cols) {
  /*
   * a, b and c are all 2d matricies
   * if you pass in 512, they will be double [512][512]
   * in this case, a_rows, a_cols and b_cols, will all be N (so if you call with
   * 512, all will be 512
   */

  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      // this will effectively set every item in C equal to 0
      c[i][j] = 0;
      for (int k = 0; k < a_cols; k++) c[i][j] += a[i][k] * b[k][j];
    }
  }
}
