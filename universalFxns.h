// Includes all universal functions for any project Euler Questions
// Sean Kane
#ifndef UNIVERSALFXNS_H
#define UNIVERSALFXNS_H
#include <cmath>
bool palindrome(int num);
bool isPrime(int num);
int divisors(int n);
double sqrt(int n);

double sqrt(int n){
  double a = double(n);
  return pow(a,0.5);
}

int divisors(int n){
  int count = 0;
  for(int a = 1; a < n; a++){
    if(n % a == 0){
      count++;
    }
  }
  return count;
}

bool palindrome(int num){
  int n, digit, rev = 0;
  n = num;
  do{
    digit = num%10;
    rev = (rev*10) + digit;
    num = num/10;
  }while(num != 0);
  if(n == rev){
    return true;
  }
  return false;
}

bool isPrime(int num){

  if(num <= 3){
    if(num == 2 || num == 3){
      return true;
    }
    return false;
  }
  if(num % 2 == 0){
    return false;
  }
  unsigned long k = int(sqrt(n)) + 1
  for(int a = 3; a < num; a++){
    if(num % a == 0){
      return false;
    }
  }
  return true;
}




#endif
