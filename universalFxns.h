// Includes all universal functions for any project Euler Questions
// Sean Kane
#ifndef UNIVERSALFXNS_H
#define UNIVERSALFXNS_H

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




#endif
