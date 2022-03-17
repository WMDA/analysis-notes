# MANOVA 

## What is a MANOVA?

A MANOVA extends the ANOVA by assessing multiple dependent variables simultaneously. A MANOVA has one one or more factors (indepedent variable such as teaching method, BMI etc) with two or more levels (i.e groups) and two or more depent variables.The advantage of a MANOVA is it can detect patterns between multiple dependent variables and independent variables.

 
## Technical details for MANOVA

More technically: 
- A MANOVA is concerned with testing if ***vectors*** of means for two or more groups are sampled from the same sampling distribution of means. 
- The null hypothsis is that that the vector of means are taken from the null distribution. 
- The alterantaive hypothesis is that one or more of the vectors are taken from a different multivariate (as there are multiple dependent varaibles) distrbution.
- A t-test and ANOVA are concerned if two or more ***means*** (not vectors) are sampled from the same sampling distribution of means (i.e from the null distribution).
- Therefore the null and alternative hypothesis are that the means are sampled from the same/different normal distribution 

### MANOVA is an ANOVA in Matrix form!

An example to work through. An experimenter has a 2x2 factor design experiment. The first factor is medicated v non-medicated, the second factor is CBT v psychotherapy. The experimenter has taken three measures, one for anxiety, one for depression and one for sleep and is interested how the groups differ on these measures following treatment. 

An ANOVA would examine the following for a single variable: A main effect for medication (do the means differ between medicated and unmedicated irrespective of psychological support), a main effect for psychology (do the means differ between CBT v psychotherapy irrepective of medication) and an interactive effect (do any of the four means differ from each other). 

A MAVOA would examine exactly the same as the ANOVA however does it for all variables at the same time. Also instead of examining means it will examine vectors. 


The variance-covariance matrix is different for ANOVAs and MANOVAs as well. While with ANOVAs it is just a single value (1x1 matrix)denoating the variance, the MANOVA has a (nxn) matrix (with ns being the number of dependent variables) shown below:  

                                      
                           depedent variable | 1 depedent variable 2 | depedent variable 3 
        -----------------------------------------------------------------------------------
        depedent variable 1    Ve1           |   cov(e1,e2)          | cov(e1,e3)
        -----------------------------------------------------------------------------------
        depedent variable 2    cov(e2,e1)    |    Ve2                |  cov(e2,e3)
        -----------------------------------------------------------------------------------  
        depedent variable 3    cov(e3,e1)    |    cov(e3,e2)         |  Ve3
        -----------------------------------------------------------------------------------


The variance (Ve) is on the diagonals and the covariance is on the off diagonals. 

How is this calculated? In an ANOVA the covariance matrix is partitioned into parts to do with error and a part to with hypothesis. Therefore the total variability is the sum of the variablity of the main effect of the factors (medication and psychology in this case), the variation of the interaction factor (medication v psychology) and the variation in the error term.
In a MANOVA the covariance matrix is calculated in exactly the same way however the variance now is a vector of the variabilty of the dependent variables. So in a MANOVA (using our example) the total variabilty equals the variability of the medication (a 3x3 matrix consinsting of the variance of each dependent variable), the variability of the psychology (a 3x3 matrix consinsting of the variance of each dependent variable), the variability of the interaction (a 3x3 matrix consinsting of the variance of each dependent variable) and varariabilty due to error (again 3x3 matrix).

``` V(total) = V(medication) + V(psychology) + V(interaction) + V(error) ``` 

**Note on the other variation matricies.** 

All the matricies should have a similar layout to the variance-covariance matrix above. The varaiation matrix for error should be treated as a *within group* covariance matrix. So this means that the off diagnoals should be treated as individuals with high scores in one factor (anxiety) within a group also tend to have high scores in another factor (depression) scores. Inspect the error matrix as if there is perfect measurement then everything should be 0. The off diagnonals in the other matricies should be treated as *between group*. So if we saw higher values in the covariance matrix for the main effect of psychology then we can assume that psychology is more efficaious on the three measures than medication. Also if the values in the psychology variation matrix are 0 then there will be no main effect of psychology. The interaction matrix has the following interpretation, if we control for medication and psychology then do groups with high average scores in depression also have high average scores in anxiety?   


### Let's go deeper bro!

An ANOVA is concerned about analysing the variance first between groups variance and then within groups (error variance). The estimates of population variance is called mean squares. 

Steps to calculate an ANOVA:

1) Calculate the total sum of squares.

```total sum of squares = Σ((y - mean(y))**2)```



2) Calculate within group variance (error or residuals).

This involves calculating the residual sum of squares then calculating the mean squared error.

Caclulating the resudial sum of squares can be done in two different ways: 

Way 1: Suming up the variance of the observed variance from each of the groups then dividing by the number of groups. The central limit theroum states that when scores are normally distrubted with a mean and variance (σ2), then the sampling distribution of means based on the size of the number of groups will also be sampled from the same normal distribution. Basically the 4 means are normally distributed with a mean and a varaince of σ2/number of groups.

``` residual sum of squares = Σ(σ2 for each grop) ```

or

Way 2: 

```rss = Σ((y - y_hat)**2)```

Calculating the mean square errors of the residuals is done by dividng the residual sum of squares by the degrees of freedom for residuals (the number of groups - 1).



3) Calculate between groups variance (froup or hypothesis).

This concerns, unlike the within group, the means of each dependent group rather than the variance of each group, It is calculated by first getting the explained sum of squares then dividing by the model degrees of freedom (number of data points -1). This can be done in two different ways:

Way 1

 ``` total sum of squares * the number in each group  ```

or way 2  ``` total sum of squares -  residual sum of squares```


An F statistic (mean squares of group/mean squares of error) or an A statistic (sum of squares of hypothesis/sum of squares of error * df of model/df of residuals) can then be calculated along with a p value 


The MANOVA is exactly as the ANOVA, it is also concerned with variance between and within groups however the only difference when calculating the within/between group variance is the equations are written in matrix form! This is to take into acount that the variance estimates are now covariance matricies (note it is easiest to calculate the within group covariance matricies by summing them all up). Also another difference when calculating the between group variance the assumption is that the normal distribution if for the vector of means rather than an individual value. 

 The difference between ANOVAs and MANVOS come working out F/A statistics, as this leads to multiple test statisitics each with their own (sometimes differnt) p values.

### Multiple Test Statistics 

For ANOVAS an F value is just a ratio and an A statistic is just a further simplified ratio (as it saves having to calculate mean squared values as everything can be calculated by just sum of squares and degress of freedom). 

These are just expressed in matrix form in the MANOVA. In MANOVS the a statistic is the important statistic and is expressed as 

```
A = HE-1

H = hypothesis sum of squares and cross products matrix
E = Error sum of squares and cross product matrix. 

Also note that as nothe H and E are symetrical that

HE-1 = E-1H
```


All significance tests for the MANOVA are done on the A statistic. 

However, there are four different multivariate tests done that are done to calculcate the A statistic.  

1) Pillai's trace. Conisdered the most powerful and robust. 
2) Hotelling-Lawley's trace.
3) Wilk's lambda (the first statistic to be used in multivariate analysis).
4) Roy's largest root. Based on the maximum eigenvalue of A=HE-1. Often considered the weakest.

Once the statistics are done they are translated into F statistics to test the null hypothesis, this is done as there are many published tables of F distribution so easy to calculate. In some cases the F statistic is sometimes approximated or exact.


## When a MANOVA is used.

1) When several correlated variables and there is a need for an overall statistical test with one value on this set of variables instead of performing multiple tests.

2) To explore how independent variables influence some patten of response in the dependent variables.

3) Used when assumptions of sphericity is violated (how equal the variance is, more equal variance = Sphericity)

## Assumptions of MANOVA

1) Response variables are continous.
2) The residuals follow the multivariate-normal probability distribution with means equal to zero.
3) The variance-covariance matrix of each group of residuals are equal.
4) Individuals are indepedent. 
5) Linear relationship among the depedent variables in each group.

### MANOVA normalcy and outliers.

MANOVA is robust to a modest amount of skness in the data (deviation from the normal distribution), if a sample size that produces 20 degrees of freedom in the model is used. However outliers causing non-normality can cause a big problem that robustness cannot overcome. So should screen data for outlyers. 

### Homogeneity of Covariance matricies

The MANOVA makes the assumption that group covariance matricies are equal. If the design is balanced this will be guaranteed due to the robustness of the MANOVA. If not then use Box's M test and if this testt if significant then this can distort pvalues. If this assumption is violated then only pillai's trace can be used.

### Multicollinearity

An issue in MANOVAs. 

## Limiations of MANOVA

The MANOVA gives one overall test of the equality of mean vectors for several groups, but like an ANOVA it cannot tell you which groups differ from which other groups on their mean vectors. 
MANOVA will also not tell you which variables are responsible for the differences in mean vectors, again it is possible to overcome this with proper contrast coding for the dependent variables.

## References

http://ibgwww.colorado.edu/~carey/p7291dir/handouts/manova1.pdf
https://www.st-andrews.ac.uk/media/ceed/students/mathssupport/ANOVA.pdf