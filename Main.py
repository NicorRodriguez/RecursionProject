from datetime import datetime

# Función que utliza la recursio para calcular la siguiente sucesión:
# n -> n/2 (n es par)
# n -> 3n + 1 (n es impar)
def suc(n):
    i = 1

    if(n != 1 and n % 2 == 0):
        i = suc(int(n/2)) + 1
    elif (n != 1):
        i = suc(int(n * 3 +1)) + 1

    return i

total = 0
max = 2

t1 = datetime.now()
for i in range(2, 1000000):
    mysuc = suc(i)
    if(mysuc > total):
        total = mysuc
        max = i
t2 = datetime.now()
delta = t2 - t1

print(delta.total_seconds()*1000, "ms")
print("Cantidad de digitos:", total)
print("Numero en cuestion::", max)