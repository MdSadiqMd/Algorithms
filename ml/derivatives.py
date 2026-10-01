# The derivative f'(x) measures the instantaneous rate of change of f with respect to x.
# Common rules:
#   d/dx x^n   = n * x^(n-1)
#   d/dx e^x   = e^x
#   d/dx ln(x) = 1/x


def derivative_power_rule(n, x):
    """Computes the derivative of x^n with respect to x."""
    return n * (x ** (n - 1))


def derivative_exp(x):
    """Computes the derivative of e^x with respect to x."""
    from math import exp

    return exp(x)


def derivative_ln(x):
    """Computes the derivative of ln(x) with respect to x."""
    return 1 / x


def gradient(f, vars, epsilon=1e-8):
    """
    Numerically computes the gradient vector of function f at the point given by vars.
    Args:
        f: function from R^n to R
        vars: list or np.array of length n
        epsilon: small number for finite difference
    Returns:
        gradient: list of partial derivatives (approximate)
    """
    grad = []
    for i, x in enumerate(vars):
        vars_forward = list(vars)
        vars_backward = list(vars)
        vars_forward[i] += epsilon
        vars_backward[i] -= epsilon
        df = f(*vars_forward) - f(*vars_backward)
        grad.append(df / (2 * epsilon))
    return grad
