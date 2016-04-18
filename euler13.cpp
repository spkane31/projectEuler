#include <iostream>
#include <cmath>
#include <ctime>
#include <fstream>
#include "universalFxns.h"

using namespace std;

void euler13ArrayCheck(int arr[], int length);
void arrayPrint(int arr[], int length);

void arrayPrint(int arr[], int length){
  for(int a = 0; a < length; a++){
    cout << arr[a] << " ";
  }
  cout << endl << endl;;
}

int main(){
  string STRING;
  int arr[100][50] = {0};
  ifstream infile;
  int x = 0;
  infile.open ("euler13input.txt");
  while(getline(infile, STRING)){
    for(int a = 0; a < 50; a++){
      arr[x][a] = (int)STRING[a] - 48;
    }
    x++;
  }

  int answer[150] = {0};
  int count = 149;
  for(int a = 49; a >= 0; a--){
    for(int b = 0; b < 100; b++){
      answer[count] += arr[b][a];
    }
    count--;
  }

  euler13ArrayCheck(answer, 150);
  arrayPrint(answer, 150);
  return 0;
}

void euler13ArrayCheck(int arr[], int length){
  for(int a = length-1; a > 0; a--){
    if(arr[a] > 10){
      int stay = arr[a] % 10;
      int move = (arr[a] - stay)/10;
      arr[a-1] += move;
      arr[a] = stay;
    }
  }
}
