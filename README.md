# uniqlist
A list contains only unique elements

## Usage
```
l := uniqlist.New()
l.PushBack(1)         /// [ 1 ]
l.PushBack("foo")     /// [ 1, "foo" ]
l.PushFront("front")  /// [ "front", 1, "foo"]
l.Remove(1)           /// [ "front", "foo"]
a := l.PopBack()      /// a == "front". List is now ["foo"]
```

## Traverse
```
for elem := l.List.Front(); elem != nil; elem = elem.Next() {
    var v interface{}
    v = elem.Value
}
```

## Goroutine Safe
Traverse list is not goroutine safe.

Goroutine safe functions:
```
New()
PushBack()
PushFront()
PopBack()
PopFront()
Remove()
Len()
```
