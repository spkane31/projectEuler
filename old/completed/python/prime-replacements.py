# import math

# def primeReplacements(n) -> int:
#   string = str(n)

#   for i in range(len(string)-1):
#     for j in range(len(string)):
#       if i == j:
#         break
#       else:
#         s = str(n)
#         s = s[:i] + "*" + s[i+1:]
#         s = s[:j] + "*" + s[j+1:]
#         count, m = replaceStars(s)
#         if count > 5:
#           print(count, m)
#           if count == 8:
#             quit()


# def replaceStars(s):
#   l = []
#   for i in range(10):
#     new = s.replace("*", str(i))
#     l.append(new)
#   return countPrimes(l)

# def countPrimes(lst):
#   count = 0
#   m = 9999999
#   for n in lst:
#     if isPrime(n):
#       count += 1
#       if int(n) < m:
#         m = int(n)
#   return count, m

# def isPrime(n):
#   num = int(n)

#   if num > 1:
#     for i in range(2, num//2):
#       if (num % i) == 0:
#         return False
#     return True
#   return False

# # print(primeReplacements(56003))
# start = 56003
# while True:
#   if isPrime(start):
#     # print(start)
#     primeReplacements(start)
#   start += 1

from collections import Counter
import time

start = time.time()
def sieve(n):
    is_prime = [True]*n
    is_prime[0] = False
    is_prime[1] = False
    is_prime[2] = True
    # even numbers except 2 have been eliminated
    for i in xrange(3, int(n**0.5+1), 2):
        index = i*2
        while index < n:
            is_prime[index] = False
            index = index+i
    prime = [2]
    for i in xrange(3, n, 2):
        if is_prime[i]:
            prime.append(i)
    return prime

# primes upto 1 million
primes = sieve(1000000)

primes = [x for x in primes if len(str(x)) - len(set(str(x))) >= 3]

def pdr(s):
  s = str(s)
  sol = []
  for duplicate in (Counter(s) - Counter(set(s))):
    a = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
    temp = [int(s.replace(duplicate, x)) for x in a]
    sol.append(temp)
  return sol

checked = []


def check(l):
    """take a list and remove all the values which are
    not prime numbers, finally return a list with only
    prime numbers"""
    for i in l:
        checked.append(i)
        if i not in primes:
            l.remove(i)
    return l

# condition for while loop
flag = True

# while loop iterator
i = 0

# while loop
while flag:
    # check if we have not check the number before
    if primes[i] not in checked:
        # find the family of the given number
        replacements = pdr(primes[i])
        for j in replacements:
            if len(check(j)) == 8:
                print j[0]
                flag = False
                break
    i += 1

# time at the end of program execution
end = time.time()

# total time taken for execution
print end - start