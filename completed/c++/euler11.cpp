#include <iostream>
#include <cmath>
#include <ctime>
#include <fstream>
#include "universalFxns.h"

using namespace std;

int testHor(int grid[20][20]);
int testVert(int grid[20][20]);
int testDiagA(int grid[20][20]);
int testDiagB(int grid[20][20]);

int main(){
  string STRING;
  int arr[20][20] = {0};
  ifstream infile;
  int x = 0, b = 0;
  infile.open ("euler11.txt");
  while(getline(infile, STRING)){
    for(int a = 0; a < 59; a += 3){
      arr[x][b] = ((int)STRING[a] - 48)*10 + ((int)STRING[a+1] - 48);
      b++;
    }
    x++;
    b = 0;
  }
  int hor, vert, DiagA, DiagB;
  hor = testHor(arr);
  vert = testVert(arr);
  DiagA = testDiagA(arr);
  DiagB = testDiagB(arr);

  cout << hor << " " << vert << " " << DiagA << " " << DiagB << endl;


  return 0;

}

int testHor(int grid[20][20]){
  int temp = 0, largest = 0;
  for(int rows = 0; rows < 20; rows++){
    for(int cols = 0; cols < 16; cols++){
      temp = 1;
      for(int x = 0; x < 4; x++){
        temp *= grid[rows][cols + x];
      }
      if(temp > largest){
        largest = temp;
      }
    }
  }
  return largest;
}

int testVert(int grid[20][20]){
  int temp = 0, largest = 0;
  for(int cols = 0; cols < 20; cols++){
    for(int rows = 0; rows < 16; rows++){
      temp = 1;
      for(int x = 0; x < 4; x++){
        temp *= grid[rows + x][cols];
      }
      if(temp > largest){
        largest = temp;
      }
    }
  }
  return largest;
}

int testDiagA(int grid[20][20]){
  int temp = 0, largest = 0;
  for(int rows = 0; rows < 16; rows++){
    for(int cols = 0; cols < 16; cols++){
      temp = 1;
      for(int x = 0; x < 4; x++){
        temp *= grid[rows + x][cols + x];
      }
      if(temp > largest){
        largest = temp;
      }
    }
  }
  return largest;
}

int testDiagB(int grid[20][20]){
  int temp = 0, largest = 0;
  for(int rows = 0; rows < 16; rows++){
    for(int cols = 3; cols < 20; cols++){
      temp = 1;
      for(int x = 0; x < 4; x++){
        temp *= grid[rows + x][cols - x];
      }
      if(temp > largest){
        largest = temp;
      }
    }
  }
  return largest;

}
