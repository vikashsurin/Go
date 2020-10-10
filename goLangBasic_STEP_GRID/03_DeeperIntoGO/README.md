### Custom type declaration 
- 


### Type conversion 
```
  []byte("string")
```

- byte slice can only convert strings
- from byte to string 
```
    string(bs)
```

### create slice of string from a string
```
    str := "apple,google,laptop,magazine"

    sliceOfString := strings.Split(str,",")
```