from collections import deque


def canFinish(numCourses: int, prerequisites: list[list[int]]) -> bool:
    sorted_order = []
    if numCourses == 0:
        return True

    # create indegree list (how many vertices have edges pointed to them)
    indegree = {i: 0 for i in range(numCourses)}
    # create empty graph - each vertex and the vertices they point to
    graph = {i: [] for i in range(numCourses)}

    # fill in the graph and indegreee
    for dep in prerequisites:
        parent, child = dep[1], dep[0]
        graph[parent].append(child)
        indegree[child] += 1

    # sources == vertices with nothing pointing to them, we start here
    sources = deque()
    for key in indegree:
        if indegree[key] == 0:
            sources.append(key)

    while sources:
        vertex = sources.popleft()
        sorted_order.append(vertex)
        # each time we get a source, we decrement the number of items pointing to its children
        for child in graph[vertex]:
            indegree[child] -= 1

            # any new sources to the list
            if indegree[child] == 0:
                sources.append(child)

    return len(sorted_order) == numCourses
