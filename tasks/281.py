# zigzag iterator but for undifined numbers of vecs

class ZigzagIterator2:
    """
    @param: vecs: a list of 1d vectors
    """

    def __init__(self, vecs):
        # do intialization if necessary
        self.vecs = vecs
        self.queue = []
        for i in range(len(vecs)):
            if len(vecs[i]) > 0:
                self.queue.append((i, 0))
    """
    @return: An integer
    """

    def _next(self):
        row, col = self.queue.pop(0)

        if col + 1 < len(self.vecs[row]):
            self.queue.append((row, col+1))
        return self.vecs[row][col]
    """
    @return: True if has next
    """

    def hasNext(self):
        return len(self.queue) != 0
