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

## Benchmarks

```
BenchmarkPushBack-12                 	 2082919	       578 ns/op
BenchmarkPushFront-12                	 2275348	       543 ns/op
BenchmarkPushBackIgnore-12           	 2297386	       522 ns/op
BenchmarkPopBack-12                  	 6493636	       198 ns/op
BenchmarkPopFront-12                 	 6493476	       203 ns/op
BenchmarkRemove-12                   	 5808315	       223 ns/op
BenchmarkContainerListPushBack-12    	 8317676	       125 ns/op
BenchmarkContainerListRemove-12      	135425845	        10.3 ns/op
BenchmarkMapPush-12                  	 3209070	       340 ns/op
BenchmarkMapRemove-12                	 8979912	       166 ns/op
```
