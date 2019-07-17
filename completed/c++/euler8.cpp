// Sean Kane
// Started: April 12, 2016
// Finished: April 12, 2016

#include <iostream>
#include <cmath>
#include "universalFxns.h"

using namespace std;

int main(){
  long long int num = 2;
  for(int a = 3; a < 2000000; a++ ){
    if(isPrime(a)){
      num += a;
      cout << "Sum: " << num << endl;
      cout << "Prime: " << a << endl;
    }
  }
  return 0;
}
