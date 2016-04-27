// Euler 18: Maximum Path Sum 1
// Sean Kane
// April 22, 2016

#include <iostream>
#include <cmath>
#include <ctime>
#include <fstream>
#include "universalFxns.h"

using namespace std;

void checkArr(int arr[], int length);

int main(){
  string STRING;
  int arr[15][15] = {0};
  ifstream infile;
  int x = 0;
  int line = 1, test = 0, other = 0;
  infile.open ("euler18input.txt");
  while(getline(infile, STRING) && line < 16){
    int col = 0;
    for(int a = 0; a < line*2 + line - 1; a+=3){
      arr[line-1][col] = (((int)STRING[a] - 48)*10) + ((int)STRING[a+1]-48);
      col++;
    }line++;
  }
  int index[14] = {0};
  int start = 75;
  int highest = 0, temp = 0;
  while(index[0] != 2){ // run until every index is GO RIGHT
    temp = start;
    for(int a = 0; a < 14; a++){ // runs for each iteration.
      temp += arr[a+1][sumArray(index,a+1)];
    }
    if(temp > highest){
      highest = temp;
      cout << "New High: " << highest << endl;
      for(int a = 0; a < 14; a++){
        cout << index[a] << " ";
      }
    }
    index[13]++;
    checkArr(index, 14);
  }

  return 0;
}

void checkArr(int arr[], int length){
  for(int a = length - 1; a > 0; a--){
    if(arr[a] == 2){
      arr[a] = 0;
      arr[a-1]++;
    }
  }
}
