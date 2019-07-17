// Sean Kane
// Project Euler Problem  4
/*
** A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.

** Find the largest palindrome made from the product of 3 digit numbers.

*/
#include <iostream>
#include <cmath>
#include "universalFxns.h"

bool palindrome(int num);

using namespace std;

int main() {
  int largest = 0;
  for(int a = 100; a < 1000; a++){
    for(int b = 100; b < 1000; b++){
      if(palindrome(a*b)){
        if(a*b > largest){
          largest = a*b;
          cout << largest << endl;
        }
      }
    }
  }
  return 0;
}
