import java.util.Random;

public class MatrixMultiplication {

  public static double[][] multiplicar(double[][] A, double[][] B) {

    int aRows = A.length;
    int aColumns = A[0].length;
    int bRows = B.length;
    int bColumns = B[0].length;

    if (aColumns != bRows) {
      throw new IllegalArgumentException("A:Rows: " + aColumns + " did not match B:Columns " + bRows + ".");
    }

    double[][] C = new double[aRows][bColumns];

    for (int i = 0; i < aRows; i++) { // aRow
      for (int j = 0; j < bColumns; j++) { // bColumn
        for (int k = 0; k < aColumns; k++) { // aColumn
          C[i][j] += A[i][k] * B[k][j];
        }
      }
    }

    return C;
  }

  public static void main(String[] args) {

    MatrixBuilder mb = new MatrixBuilder();
    int n = 100;
    if(args.length > 0) {
      n = Integer.parseInt(args[0]);
    }
    double[][] A = mb.setX(n).setY(n).enableGenRand().build();
    double[][] B = mb.setX(n).setY(n).enableGenRand().build();

    double[][] result = multiplicar(A, B);
  }
}


class MatrixBuilder {

  int x = 0;
  int y = 0;
  boolean genRand = false;

  public MatrixBuilder(){
  }

  public MatrixBuilder setX(int x){
    this.x = x;
    return this;
  }

  public MatrixBuilder setY(int y){
    this.y = y;
    return this;
  }

  public MatrixBuilder enableGenRand(){
    this.genRand = true;
    return this;
  }

  public double[][] build(){
    if (x <= 0 || y <= 0){
      return null;
    }

    double[][] res = new double[x][y];

    if (genRand) {
      Random rand = new Random();

      for (int i = 0; i < x; i++) {
        for (int j = 0; j < y; j++) {
          res[i][j] = rand.nextDouble();
        }
      }
    }

    return res;
  }
}
