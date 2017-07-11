## Golang Reflection Example with Custom Types and Nested Slices

Get a custom struct type's field by name
```
r := reflect.ValueOf(&yourStructType)
f := reflect.Indirect(r).FieldByName("fieldNameInYourStructType")
```

Convert that return Value into it's proper type
```
yourStructTypes, ok := f.Interface().([]YourStructType)
if !ok {
    panic("f is not a []YourStructType")
}
```
