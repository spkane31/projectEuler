// Sean Kane
// Euler Problem 15
// Started: April 19, 2016
// Finished:

#include <iostream>
#include <cmath>
#include "universalFxns.h"

using namespace std;

int main(){
  unsigned long long int grid[21][21] = {0};
  grid[0][0] = 1;
  for(int a = 1; a < 21; a++){
    grid[0][a] = 1;
    grid[a][0] = 1;
  }
  for(int b = 1; b < 21; b++){
    for(int c = 1; c < 21; c++){
      grid[b][c] += grid[b-1][c] + grid[b][c-1];
    }
  }
  cout << grid[2][1] << endl;
  cout << grid[3][2] << endl;
  cout << grid[4][4] << endl;
  cout << grid[20][20] << endl;

  for(int b = 0; b < 21; b++){
    for(int c = 0; c < 21; c++){
      cout << grid[b][c] << " ";
    }
    cout << "\n";
  }
  return 0;
}
