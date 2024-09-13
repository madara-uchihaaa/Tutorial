package calculator;
class VeryAdvanceCalculator extends AdvanceCalculator{
    public int square(int x){
        return x*x;
    }
    public int cube(int x){
        return x*x*x;
    }
    public double log(double x){
        return Math.log(x);
    }
    public double log10(double x){
        return Math.log10(x);
    }
}

// AdvanceCalculator.java -> VeryAdvanceCalculator.java // Single Inheritance


// Path: VeryVeryAdvanceCalculator.java -> VeryAdvanceCalculator.java -> AdvanceCalculator.java -> Calculator.java

// Multilevel Inheritance