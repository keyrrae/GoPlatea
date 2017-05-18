def mergeKLists(self, lists):
    if not lists:
        return None
        
    sentinel = ListNode('0')
    while len(lists) > 1:
        merged = []
        while len(lists) > 1:
            merged.append(self.merge(lists.pop(), lists.pop(), sentinel))
        lists += merged
    return lists[0]
    
    
def merge(self, x, y, s):
    current = s
    while x and y:
        if x.val < y.val:
            current.next = x
            x = x.next
        else:
            current.next = y
            y = y.next
        current = current.next
    current.next = x if x else y
    return s.next