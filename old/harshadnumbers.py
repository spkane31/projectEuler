import math, time
import numpy as np

globalTruncs = []

def SumDigits(n):
  s = n % 10
  n = int(n/10)
  while n != 0:
    s += n % 10
    n = int(n/10)
    
  return s

def IsHarshad(n):
  return n % SumDigits(n) == 0

def IsRightTH(n):
  while n != 0:
    if not IsHarshad(n):
      return False
    else:
      n = int(n/10)
  return True

def IsStrongHarshad(n):
  if not IsHarshad(n):
    return False
  
  n = int(n / SumDigits(n))

  return IsPrime(n)

def IsPrime(n):
  if n < 2: return False
  if n == 2: return True
  i = 2
  while i < int(math.sqrt(n)+1):
    if n % i == 0:
      return False
    i += 1
  return True

def IsStrongRTH(n):
  if not IsPrime(n):
    return False

  n = int(n/10)
  if n < 10:
    return False
  

  return IsStrongHarshad(n) and IsRightTH(n)

def GenForN(n):
  if n > 10e13:
    return []
  ret = []

  for j in range(10):
    if IsHarshad(n * 10 + j):
      globalTruncs.append(n * 10 + j)
      # print(n*10+j)
      GenForN(n * 10 + j)


  return ret

def GenerateRightTruncatable():
  LIMIT = 10e13
  print(LIMIT)


  start = [1, 2, 3, 4, 5, 6, 7, 8, 9]

  ret = []

  for each in start:
    if IsHarshad(each):
      globalTruncs.append(each)
      GenForN(each)
  # print(globalTruncs)
  print(len(globalTruncs))

  s = 0

  for each in globalTruncs:
    n = each / SumDigits(each)
    if IsPrime(n):
      l = FindSHRTNums(each)
      if len(l):
        for i in l:
          if i < LIMIT:
            s += i
            # print(i)
            print(s)

def FindSHRTNums(n):
  ret = []
  for i in range(10):
    if IsPrime(n * 10 + i):
      ret.append(n*10+i)
  return ret
  


def main():
  # a = np.array([1, [2, 3], [3, [4, 5]]])
  # print(a.flatten())
  
  t = time.time()
  GenerateRightTruncatable()
  # quit()


  # # print(IsHarshad(201))
  # # print(IsRightTH(201))
  # # print(IsStrongHarshad(201))
  # # print(IsStrongRTH(2011))


  # SRTH = []
  # i = 3
  # while i < 10e14: #int(10e14):
  #   # if IsStrongRTH(i):
  #   if IsStrongHarshad(i) and IsRightTH(i):
  #     if IsPrime(i):
  #       SRTH.append(i)
  #       print(i)
  #     for j in range(1, 10, 2):
  #       if IsPrime(i * 10 + j):
  #         SRTH.append(i * 10 + j)
  #         print(i * 10 + j)
  #         # print(i, i * 10 + j)
  #   i+=1
  # print(sum(SRTH))
  # print(f"{time.time() - t}\n\n")

  # t = time.time()
  # s = 0
  # i = 2
  # while i < 10000:
  #   if IsStrongRTH(i):
  #     s += i
  #     # print(i)
  #   i += 1
  # print(s)
  print(f"{time.time() - t}\n\n")
  # print("done")
  

main()