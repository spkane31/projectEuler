// Sean Kane
// Started: April 12, 2016
// Finished:
// Euler Problem 14

#include <iostream>
#include <cmath>
#include "universalFxns.h"

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

using namespace std;

int main(){
  cout << "Hello" << endl;;
  unsigned int a, iter, highest = 0, num = 0;
  for(unsigned int i = 3; i < 1000000 ; i++){
    a = i;
    iter = 1;
    while(a > 1){
      if((a % 2) == 0){
        a = a / 2;
      }else{
        a = (3 * a) + 1;
      }
      iter++;
    }
    if(iter > highest) {
      highest = iter;
      num = i;
    }
  }
  cout << "\n\n";
  cout << "Answer: ";
  cout << highest << " " << num << endl;
  return 0;
}
