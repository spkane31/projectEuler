// Sean Kane
// Started: April 20, 2016
// Finished: April 20, 2016
// Time: .122 seconds
/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
*/


#include <iostream>
#include <cmath>
#include <ctime>
#include "universalFxns.h"

using namespace std;

int main(){
  int array[1001] = {0};
  int ones = 0, tens = 0, hundreds = 0;
  array[1] = 3;
  array[2] = 3;
  array[3] = 5;
  array[4] = 4;
  array[5] = 4;
  array[6] = 3;
  array[7] = 5;
  array[8] = 5;
  array[9] = 4;
  array[10] = 3;
  array[11] = 6;
  array[12] = 6;
  array[13] = 8;
  array[15] = 7;
  array[18] = 8;
  array[20] = 6;
  array[30] = 6;
  array[40] = 5;
  array[50] = 5;
  array[60] = 5;
  array[70] = 7;
  array[80] = 6;
  array[90] = 6;
  array[1000] = 11;
  for(int a = 14; a <= 19; a++){
    if(a == 15 || a == 18){
      a++;
    }
    array[a] = array[a % 10] + 4;
  }
  for(int a = 21; a < 100; a++){
    if(array[a] != 0){}else{
      tens = a/10;
      ones = a - tens*10;
      array[a] = array[tens*10] + array[ones];
    }
  }

  for(int a = 100; a < 1000; a++){
    if(a % 100 == 0){
      hundreds = a/100;
      array[a] = array[hundreds] + 7;
    }else{
      hundreds = a/100;
      tens = (a - hundreds*100);
      array[a] = array[hundreds] + 7 + array[tens] + 3; //"hundred" and "and"
    }
  }
  unsigned long long int answer = sumArray(array, 1001);
  cout << "Result: " << answer << endl;
  return 0;
}
