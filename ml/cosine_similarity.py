# Cosine similarity measures alignment: cos(u, v) = (u · v) / (||u||2 * ||v||2)
# 1 means same direction, 0 means orthogonal, -1 means opposite directions

import math
from collections.abc import Sequence

Number = int | float
Vector = Sequence[Number]

u: Vector = [1, 2, 3]
v: Vector = [4, 5, 6]


def dot(u: Vector, v: Vector) -> Number:
    if not u or not v:
        raise ValueError("vectors must be non-empty")
    if len(u) != len(v):
        raise ValueError("vectors must have the same length for dot product")
    total: Number = 0
    for i in range(len(u)):
        total += u[i] * v[i]
    return total


def l2_norm(v: Vector) -> float:
    if not v:
        raise ValueError("vector must be non-empty")
    total = 0
    for i in range(len(v)):
        total += v[i] ** 2
    return math.sqrt(total)


def cosine_similarity(u: Vector, v: Vector) -> float:
    if not u or not v:
        raise ValueError("vectors must be non-empty")
    if len(u) != len(v):
        raise ValueError("vectors must have the same length for cosine similarity")
    norm_u = l2_norm(u)
    norm_v = l2_norm(v)
    if norm_u == 0 or norm_v == 0:
        raise ValueError("cosine similarity is undefined for zero vectors")
    return dot(u, v) / (norm_u * norm_v)


print(cosine_similarity(u, v))
