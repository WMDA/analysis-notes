'''
Important: this script will keep running indefinetly unless closed manually.
'''

import rpy2.situation
import rpy2.robjects as robjects
import rpy2.robjects.packages as rpackages
 
r = robjects.r

x_values = robjects.IntVector(range(10))
y = r.rnorm(10)
r.X11()
r.layout(r.matrix(robjects.IntVector([1,2,3,4]), nrow=2, ncol=2))
r.plot(r.runif(10), y, xlab='runif', ylab='foo/bar', col='red')
while True:
    continue