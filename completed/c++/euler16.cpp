// Sean Kane
// Started: April 12, 2016
// Finished: April 14, 2016
// Euler Problem 16
/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

// Highest storable number is 2^63 = 9223372036854775808

#include <iostream>
#include <cmath>
#include <ctime>
#include "universalFxns.h"

using namespace std;

void euler16ArrayCheck(int arr[], int length);
void arrayPrint(int arr[], int length);

void arrayPrint(int arr[], int length){
  for(int a = 0; a < length; a++){
    cout << arr[a] << " ";
  }
  cout << endl << endl;;
}

void euler16ArrayCheck(int arr[], int length){
  for(int a = 0; a < length-1; a++){
    if(arr[a] > 10){
      int stay = arr[a] % 10;
      int move = (arr[a] - stay)/10;
      arr[a+1] += move;
      arr[a] = stay;
    }
    if(arr[a] == 10){
      arr[a+1] += 1;
      arr[a] = 0;
    }
  }

}


using namespace std;

int main(){
  int start_s = clock();
  int digits[302] = {0};

  int power = 1;
  digits[0] = 2;
  for(power; power < 1000; power++){
    for(int a = 0; a < 302; a++){
      digits[a] = digits[a] * 2;
    }
    euler16ArrayCheck(digits, 302);
    // if(power % 200 == 0){
      // arrayPrint(digits, 302);
      // cout << "\n\n\n\n";
    // }
  }

  long int sum = sumArray(digits, 302);
  cout << endl << endl;
  cout << sum << endl;
  int stop_s = clock();
  cout << "Time: " << (stop_s - start_s)/double(CLOCKS_PER_SEC)*1000 << endl;
  return 0;
}
