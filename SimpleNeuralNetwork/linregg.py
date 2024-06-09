import numpy as np

x=np.array([1,2,3,4,5])
y=np.array([2,4,5,4,5])
z=np.array([0,6,7,8,9])

mean_x=np.mean(x)
mean_y=np.mean(y)

m=np.sum((x-mean_x)*(y-mean_y))/np.sum((x-mean_x)**2)
b=mean_y-m*mean_x

def predict(x):
    return m*x+b
def loss(c):
    return predict(c)-mean_y

print(predict(0))