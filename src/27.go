import matplotlib.pyplot as plt
from numpy import linspace

def draw_linear_line(x_points, y_points):
    """
    Draw a linear line plot of given x and y points.
    
    Args:
        x_points: A sequence or list of x values.
        y_points: A sequence or list of y values corresponding to each x value in x_points.
        
    Returns:
        None
    """
    plt.figure(figsize=(10, 5))
    plt.plot(x_points, y_points)
    plt.xlabel("X-axis")
    plt.ylabel("Y-axis")
    plt.title("Linear Line Plot Example")
    plt.show()

# Example usage
x = [0, 2, 4, 6]
y = [1, 3, 5, 7]

draw_linear_line(x, y)
