'''
# Sample code to perform I/O:

name = input()                  # Reading input from STDIN
print('Hi, %s.' % name)         # Writing output to STDOUT

# Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
'''

# Checks if size is valid integer
def checkNumber(size):
    try:
        size = int(size)
    except ValueError:
        print("Error")
        exit(2)
    return size

# Checks if size is greater than 0
def checkPositive(size):
    if not size > 0:
        print("Error")
        exit(3)

# Checks if there are non-equal key-value pairs to input size
def checkOverloaded(size, kv):
    if len(kv) != size:
        print("Overloaded")
        exit(4)

# Converts 
def compileIntoDict(x):
    d = {}
    for i in x:
        k, v = i.split(":")
        if k not in d.keys():
            d[k] = v
        elif type(d[k]) == str:
            d[k] = [d[k], v]
        else:
            d[k].append(v)
    return d

size = 6
testcase = "a:b,b:2,c:2,d:23,c:123, asd:2"
kv = [":".join(x.strip().split(":")[::-1]) for x in testcase.split(",")]
size = checkNumber(size)
checkPositive(size)
checkOverloaded(size, kv)
resultDict = compileIntoDict(kv)
print(resultDict)

