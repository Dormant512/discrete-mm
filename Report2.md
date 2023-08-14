# Discrete MM, practical task 2

**Student:** Grigoryev Mikhail

**Group:** J4133c

## Task description

- Perform a sensitivity analysis for a demographic model with respect to a set of parameters: fertility rate, boys/girls ratio, "survival" rate for different age groups (not all can be taken).
- Model output: number of inhabitants for a given year.
- Test on the final forecast values for 10, 20, 50, 100 years.
- Define ranges of model parameter values from data for previous periods (1950-2000).
- Based on all ranges of parameter values, perform an uncertainty analysis in the form of a graph with prediction intervals of the results. The values between the boundaries can be considered evenly distributed.

## Solution method

To solve the tasks the following equations were leveraged (same as practical task 1):
$$
\text{FertilityRate}(\text{year}) = \frac{\text{numNewborns(year)}}{\sum_{20}^{39} \text{womenPopulation}} \\
\text{SurvivalRate} = \text{PopulationVector(1...n)} \cdot \text{PopulationVector(0...n-1)}^{-1} \\
\text{BoysRatio} = \frac{\text{NewbornBoys}}{\text{NewbornBoys + NewBornGirls}}
$$
For modeling (code, same as practical task 1):

```python
for k in range(k_max):
        prof_m, prof_f = np.zeros(len(GROUPS)), np.zeros(len(GROUPS))
        prof_f[1:] = sr_f * df_prof_f[-1][:-1]
        prof_m[1:] = sr_m * df_prof_m[-1][:-1]
        prof_f[0] = fert_rate * (prof_f[n1:n2]).sum()
        prof_m[0] = fert_rate * (prof_f[n1:n2]).sum() * mf_ratio
```

Overall, the data was preprocessed in a simple manner (getting rid of empty rows) and the modeling was conducted in several simple steps.

1. Newborns were calculated as fertility rate $\times$ the total population of women able to give birth.
2. Other age groups were calculated as survival rate $\times$ the population of the younger group in previous iteration.

SALib was used as the library of choice for sensitivity and uncertainty analysis. Sobol indices were calculated automatically for estimating the sensitivity of certain parameters ("survival" rate, fertility rate and boys ratio). The calculations were conducted against age groups and sexes.

Uncertainty analysis was conducted for the female and male population in the time period of 2005-2050. 

## Results for the two sexes

The experiment was conducted on the demographic data from Russian Federation.

Sensitivity analysis:

![pic7](/home/dormant/discrete-mm/pic7.png)

Uncertainty analysis for females:

![pic9](/home/dormant/discrete-mm/pic9.png)

For males:

![pic8](/home/dormant/discrete-mm/pic8.png)

## Conclusions

In this practical work a demographic dataset was preprocessed and analyzed via simple iterative parametric modeling with features such as "survival" rate, fertility rate and boys/girls ratio as parameters. Additionally, sensitivity analysis regarding those parameters was conducted via calculating Sobol indices. Lastly, the uncertainty of prognosis of male and female populations onto the year 2050 was estimated and plotted as quantiles.