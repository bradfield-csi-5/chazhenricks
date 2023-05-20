from collections import deque

# Correct Order


def find_compilation_order(dependencies):
    graph = {}
    indegree = {}
    sorted_order = []

    # initialize graph, indegree
    for item in dependencies:
        parent, child = item[1], item[0]
        graph[parent], graph[child] = [], []
        indegree[parent], indegree[child] = 0, 0

    if len(graph) <= 0:
        return sorted_order

    # populate graph/indegree
    for dep in dependencies:
        parent, child = dep[1], dep[0]
        graph[parent].append(child)
        indegree[child] += 1

    # find sources
    sources = deque()
    for key in indegree:
        if indegree[key] == 0:
            sources.append(key)

    # start popping off 0 indegree items
    while sources:
        vertex = sources.popleft()
        sorted_order.append(vertex)

        for child in graph[vertex]:
            indegree[child] -= 1
            if indegree[child] == 0:
                sources.append(child)

    return sorted_order


order = [["B", "C"], ["C", "A"], ["A", "F"]]
print(find_compilation_order(order))
