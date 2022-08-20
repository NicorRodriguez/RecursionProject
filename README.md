# Recursion Project
---
This is a little script made in python and golang, the main goal of this project was to show off the difference in preformance between go and python when doing recursions.

## Code
---
### Python
```python
def suc(n):
    i = 1

    if(n != 1 and n % 2 == 0):
        i = suc(int(n/2)) + 1
    elif (n != 1):
        i = suc(int(n * 3 +1)) + 1

    return i
```

### Golang

```go
func suc(n int) int {
    i := 1

    if n != 1 && n % 2 == 0 {
        i = suc(int(n/2)) + 1
    } else if n != 1 {
        i = suc(int(n * 3 +1)) + 1
    }

    return i
}
```

As apreciated both codes are identical and use recursion to calculate the recursion. Of course there are better ways to solve this problem but it was out of the scope of the project.

## Testing
---
The problem that the algorithm is solving is to find the number that needs the largest sequence of numbers in the algorithm to get to 1.

Example:

    Input: 13
    Sequence: 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

In order to test the performance of the programming languaje, we are testing to calculate the sequence of numbers from **1** to **1000000** and create two timestamps, the first one at the beginning of the script and the second one once the script has finished, with this timestamps we calculated the difference in time and determined how long the script was running for.

    Python: 20033.763ms ~ 20s
    Golang: 846.631416ms ~ 0.8s

Around **20x** times faster when running in golang.

> :warning: This test was made on my personal computer therefore the    results may vary in your environment.

