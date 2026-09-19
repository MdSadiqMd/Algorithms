import math


def _percentile(sorted_vals: list, p: float) -> float:
    # Linear interpolation (same method as numpy default).
    n = len(sorted_vals)
    if n == 1:
        return float(sorted_vals[0])
    rank = (p / 100) * (n - 1)
    low = math.floor(rank)
    high = math.ceil(rank)
    if low == high:
        return float(sorted_vals[int(rank)])
    weight = rank - low
    return float(sorted_vals[low] * (1 - weight) + sorted_vals[high] * weight)


def descriptive_statistics(data: list) -> dict:
    """
    Calculate various descriptive statistics metrics for a given dataset.

    Args:
        data: List or numpy array of numerical values

    Returns:
        Dictionary containing mean, median, mode, variance, standard deviation,
        percentiles (25th, 50th, 75th), and interquartile range (IQR)
    """
    values = list(data)
    if not values:
        raise ValueError("data must be non-empty")
    n = len(values)

    total = 0
    for x in values:
        total += x
    mean = total / n

    ordered = sorted(values)
    mid = n // 2
    if n % 2 == 0:
        median = (ordered[mid - 1] + ordered[mid]) / 2
    else:
        median = float(ordered[mid])

    frequency = {}
    for x in values:
        frequency[x] = frequency.get(x, 0) + 1
    mode = max(frequency, key=lambda k: frequency[k])
    if hasattr(mode, "item"):
        try:
            mode = mode.item()
        except (AttributeError, ValueError):
            pass

    squared_diff_sum = 0
    for x in values:
        squared_diff_sum += (x - mean) ** 2
    variance = squared_diff_sum / n
    standard_deviation = math.sqrt(variance)

    p25 = _percentile(ordered, 25)
    p50 = _percentile(ordered, 50)
    p75 = _percentile(ordered, 75)
    iqr = p75 - p25

    return {
        "mean": round(float(mean), 4),
        "median": round(float(median), 4),
        "mode": mode,
        "variance": round(float(variance), 4),
        "standard_deviation": round(float(standard_deviation), 4),
        "25th_percentile": round(float(p25), 4),
        "50th_percentile": round(float(p50), 4),
        "75th_percentile": round(float(p75), 4),
        "interquartile_range": round(float(iqr), 4),
    }


if __name__ == "__main__":
    sample = [1, 2, 2, 3, 4, 4, 4, 5]
    print(descriptive_statistics(sample))
