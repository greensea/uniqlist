package uniqlist

import (
    "testing"
    "container/list"
)

func checkOrder(t *testing.T, s *UniqList, expect []int) {
    if s.List.Len() != len(expect) {
        t.Errorf("Length not match")
    }
    
    i := 0
    for e := s.List.Front(); e != nil; e = e.Next() {
        if e.Value.(int) != expect[i] {
            t.Errorf("Invalid order")
        }
        i++
    }
}


func TestCase1(t *testing.T) {
    s := New()
    s.PushBack(1)
    s.PushBack(2)
    
    checkOrder(t, s, []int{1, 2})
    
    s.PushFront(3)
    checkOrder(t, s, []int{3, 1, 2})
    
    s.PushBack(3)
    checkOrder(t, s, []int{1, 2, 3})
    
    s.PushFront(4)
    checkOrder(t, s, []int{4, 1, 2, 3})
    
    s.PushBackIgnore(4)
    checkOrder(t, s, []int{4, 1, 2, 3})
    
    s.PushBackIgnore(2)
    checkOrder(t, s, []int{4, 1, 2, 3})
    
    s.PushFrontIgnore(2)
    checkOrder(t, s, []int{4, 1, 2, 3})
    
    if s.Len() != 4 {
        t.Error("Invalid length")
    }
    
    
    s.Remove(1)
    checkOrder(t, s, []int{4, 2, 3})
    
    s.Remove(404)
    checkOrder(t, s, []int{4, 2, 3})
    
    
    e1 := s.PopFront()
    if e1.(int) != 4 {
        t.Errorf("Expect 4, but got %d", e1.(int))
    }
    
    e2 := s.PopBack()
    if e2.(int) != 3 {
        t.Errorf("Expect 3, but got %d", e2.(int))
    }
    
    s.PushBackIgnore(10)
    checkOrder(t, s, []int{2, 10})
    
    
    s.Remove(404)
}


func BenchmarkPushBack(b *testing.B){
    s := New()
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        s.PushBack(i)
    }
}

func BenchmarkPushFront(b *testing.B){
    s := New()
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        s.PushFront(i)
    }
}

func BenchmarkPushBackIgnore(b *testing.B){
    s := New()
    
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        s.PushBackIgnore(i)
    }
}

func BenchmarkPopBack(b *testing.B) {
    s := New()
    
    for i:=0;i<b.N;i++{
        s.PushBackIgnore(i)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.PopFront()
    }
}

func BenchmarkPopFront(b *testing.B) {
    s := New()
    
    for i:=0;i<b.N;i++{
        s.PushFrontIgnore(i)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.PopFront()
    }
}

func BenchmarkRemove(b *testing.B) {
    s := New()
    
    for i:=0;i<b.N;i++{
        s.PushFrontIgnore(i)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.Remove(i)
    }
}

func BenchmarkContainerListPushBack(b *testing.B) {
    var s list.List
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.PushBack(i)
    }
}


func BenchmarkContainerListRemove(b *testing.B) {
    var s list.List
    
    for i:=0;i<b.N;i++{
        s.PushFront(i)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.Remove(s.Front())
    }
}

func BenchmarkMapPush(b *testing.B) {
    s := make(map[int]struct{})
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s[i] = struct{}{}
    }
}


func BenchmarkMapRemove(b *testing.B) {
    s := make(map[int]struct{})
    
    for i:=0;i<b.N;i++{
        s[i] = struct{}{}
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        delete(s, i)
    }
}
