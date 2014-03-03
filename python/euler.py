primes = [2]

def prime(n):
    if len(primes) >= n:
        return primes[n-1]
    while len(primes) < n:
        start = primes[-1] + 1
        while not primeQ(start):
            start += 1
        primes.append(start)
    return prime(n)

#speed up using prime list?
def factor(n):
    if n <= 2:
        return [n]
    test = 2
    while test**2 <= n:
        if n%test == 0:
            sub = factor(n/test)
            sub.append(test)
            return sub
        test += 1
    return [n]

def primeQ(n):
    if len(factor(n)) == 1:
        return True
    return False

#Euclid's algorithm
def GCD(a,b):
    while b !=0:
        t = b
        b = a%b
        a = t  
    return a

def LCM(a,b):
    return a*b / GCD(a,b)

#######################
###Special Sequences###
#######################
fibCache = dict({0:1, 1:1})

def fib(n):
    if n in fibCache:
        return fibCache[n]
    fibCache[n] = fib(n-1) + fib(n-2)
    return fib(n)

def palindrome(n):
    #make sure we've got string
    word = str(n)

    for i in range(len(word)/2):
        if word[i] != word[-1-i]:
            return False

    return True

