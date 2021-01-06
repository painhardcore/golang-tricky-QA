# Answer
```
2
1
3
```
# Explanation

1. `1` goroutine fired at first, but it's blocks instantly for 2 second with `sleep`. 
2. `2` gorutine will be executed ASAP, but with no guarantee(program can exit before).
3. `wg.Wait()` blocks execution of the `main()` execution(obviously we'll wait 2s+ for `1` goroutine). Here execution will be switched to `2` goroutine, and "2" will be printed.
4. After 2 seconds will print "1" from `1` goroutine, and we unblock the `main()`
5. "3" will printed then.
