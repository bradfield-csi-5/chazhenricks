import operator

OPERATORS = {
    "+": operator.add,
    "-": operator.sub,
    "*": operator.mul,
    "/": operator.truediv,
}

LEFT_PAREN = "("
RIGHT_PAREN = ")"


def build_parse_tree(expression):
    tree = {}
    stack = [tree]
    node = tree

    for token in expression:
        if token == LEFT_PAREN:
            node["left"] = {}
            stack.append(
                node
            )  # confusing that we add the root to the stack twice at the begining, but I guess it doesnt matter
            node = node["left"]
        elif token == RIGHT_PAREN:
            node = stack.pop()
        elif token in OPERATORS:
            node["val"] = token
            node["right"] = {}
            stack.append(node)
            node = node["right"]
        else:
            node["val"] = int(token)
            parent = stack.pop()
            node = parent

    return tree


def evaluate(tree):
    try:
        operate = OPERATORS[tree["val"]]
        return operate(evaluate(tree["left"]), evaluate(tree["right"]))
    except KeyError:
        return tree["val"]


parse_tree = build_parse_tree("(3+(4*5))")
result = evaluate(parse_tree)
print(result)
