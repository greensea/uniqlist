package uniqlist

import (
    "container/list"
    "sync"
)

type UniqList struct {
    mu sync.Mutex
    Map map[interface{}]*list.Element
    List list.List
}

func New() *UniqList {
    return &UniqList{
        Map: make(map[interface{}]*list.Element),
    }
}

// Get the length of the list
func (t *UniqList) Len() int {
    return t.List.Len()
}

// Push an element to the end of the list. If the element is already in the list, it will be move to the end of the list.
func (t *UniqList) PushBack(v interface{}) interface{} {
    t.mu.Lock()
    defer t.mu.Unlock()
   
    e, ok := t.Map[v]
    if ok == true {
        t.List.Remove(e)
    }
    
    e2 := t.List.PushBack(v)
    t.Map[v] = e2

    return e2.Value
}

// Push an element to the front of the list. If the element is already in the list, it will be move to the front of the list.
func (t *UniqList) PushFront(v interface{}) {
    t.mu.Lock()
    defer t.mu.Unlock()
   
    e, ok := t.Map[v]
    if ok == true {
        t.List.Remove(e)
    }
    
    e2 := t.List.PushFront(v)
    t.Map[v] = e2

    //return e2.Value
}

// Push an element to the end of the list. If the element is already in the list, this func do nothing.
func (t *UniqList) PushBackIgnore(v interface{})  {
    t.mu.Lock()
    defer t.mu.Unlock()
   
    e, ok := t.Map[v]
    if ok == true && e != nil {
        //return e.Value
    } else {
        e2 := t.List.PushBack(v)
        t.Map[v] = e2
    }
}

// Push an element to the front of the list. If the element is already in the list, this func do nothing.
func (t *UniqList) PushFrontIgnore(v interface{})  {
    t.mu.Lock()
    defer t.mu.Unlock()
   
    e, ok := t.Map[v]
    if ok == true && e != nil {
        //return e.Value
    } else {
        e2 := t.List.PushFront(v)
        t.Map[v] = e2
    }
}

// Remove an element from the list.
func (t *UniqList) Remove(v interface{}) {
    t.mu.Lock()
    defer t.mu.Unlock()
    
    e, ok := t.Map[v]
    if ok != true {
        return
    }
    
    delete(t.Map, v)
    t.List.Remove(e)
    
    //return e.Value
}

// Get the first element and remove it from list
func (t *UniqList) PopFront() interface{} {
    t.mu.Lock()
    defer t.mu.Unlock()
    
    e := t.List.Front()
    if e == nil {
        return nil
    }
    
    delete(t.Map, e.Value)
    t.List.Remove(e)
    
    return e.Value
}

// Get the last element and remove it from list
func (t *UniqList) PopBack() interface{} {
    t.mu.Lock()
    defer t.mu.Unlock()
    
    e := t.List.Back()
    if e == nil {
        return nil
    }
    
    delete(t.Map, e.Value)
    t.List.Remove(e)
    
    return e.Value
}
