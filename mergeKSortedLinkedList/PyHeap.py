def mergeKLists(self, lists):
    current = sentinel = ListNode(0)
    lists = [(i.val, i) for i in lists if i]
    heapq.heapify(lists)
    while lists:
        current.next = heapq.heappop(lists)[1]
        current = current.next
        if current.next:
            heapq.heappush(lists, (current.next.val, current.next))
    return sentinel.next