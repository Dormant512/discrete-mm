# Discrete MM, practical task 1

**Student:** Grigoryev Mikhail

**Group:** J4133c

## Task description

- Determine the "survival" rates independently for men and women for all age groups (“0-4” $\to$ “5-9” $\to$ “10-14”...) according to 2000-2005 years (data for Russia or any other country).
- Determine the fertility rate for women in the age category “20-...-39”. 
- Calculate boys/girls ratio for newborn children.
- Predict the change in the country's population and demographic profile for 100 years and compare with existing prediction!

## Solution method

To solve the tasks the following equations were leveraged:
$$
\text{FertilityRate}(\text{year}) = \frac{\text{numNewborns(year)}}{\sum_{20}^{39} \text{womenPopulation}} \\
\text{SurvivalRate} = \text{PopulationVector(1...n)} \cdot \text{PopulationVector(0...n-1)}^{-1} \\
\text{BoysRatio} = \frac{\text{NewbornBoys}}{\text{NewbornBoys + NewBornGirls}}
$$
For modeling (code):

```python
for k in range(k_max):
        prof_m, prof_f = np.zeros(len(GROUPS)), np.zeros(len(GROUPS))
        prof_f[1:] = sr_f * df_prof_f[-1][:-1]
        prof_m[1:] = sr_m * df_prof_m[-1][:-1]
        prof_f[0] = fert_rate * (prof_f[n1:n2]).sum()
        prof_m[0] = fert_rate * (prof_f[n1:n2]).sum() * mf_ratio
```

Overall, the data was preprocessed in a simple manner (getting rid of empty rows) and the modeling was conducted in several simple steps:

1. Newborns were calculated as fertility rate $\times$ the total population of women able to give birth.
2. Other age groups were calculated as survival rate $\times$ the population of the younger group in previous iteration.

## Results

In this work not only age profiles were modeled. Two trends were extracted for the male and female populations. The experiment was conducted on the demographic data from Russian Federation.

### Results for sexes

![pic1](/home/dormant/discrete-mm/pic1.png)

Trend in estimated data for the whole population. Predictions:

![pic2](/home/dormant/discrete-mm/new1.png)

![pic3](/home/dormant/discrete-mm/new2.png)

![pic4](/home/dormant/discrete-mm/new3.png)



### Results for age groups

Example of the initial estimation:

![pic5](/home/dormant/discrete-mm/pic5.png)

Example of a prediction:

![pic6](/home/dormant/discrete-mm/pic6.png)



## Conclusions

In this practical work a demographic dataset was preprocessed and analyzed via simple iterative parametric modeling with features such as "survival" rate, fertility rate and boys/girls ratio as parameters.