# LINEAR-STATS

This project calculates the linear regression line and the Pearson correlation coefficient for given data.

Linear regression is a statistical method for modeling the relationship between a dependent variable and one or more independent variables. It is a powerful tool that can be used to make predictions and understand the relationships between variables.

The linear regression line is a straight line that best fits the data. It is calculated using the least squares method, which minimizes the sum of the squared residuals. The residuals are the differences between the actual values of the dependent variable and the predicted values of the dependent variable based on the regression line.

The Pearson correlation coefficient is a measure of the linear relationship between two variables. It can range from -1 to 1, with a value of 0 indicating no linear relationship and a value of 1 indicating a perfect positive linear relationship.

## AUTHORS

* Husain

## USAGE

First we need to allow permissions to use bash for math-skills file.
```
$chmod u+x linear-stats
$./linear-stats
```

Will generate a text file "data.txt" so we can use this file to test the pargrom by the following code.

```
$go run main.go data.txt
```
The result will shown in the terminal with the:

* The equation of the linear regression line

* The Pearson correlation coefficient

## LICENSE

This project is licensed under the Reboot License.