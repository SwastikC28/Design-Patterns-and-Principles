import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

interface IDepartment {
    int getTotalCost();
}

class Department implements IDepartment {
    String name;

    List<IDepartment> departments;

    Department(String name, List<IDepartment> departments) {
        this.name = name;
        this.departments = departments;
    }

    @Override
    public int getTotalCost() {
        System.out.println("Department : " + this.name);

        int sum = 0;
        for (IDepartment department : departments) {
            sum += department.getTotalCost();
        }

        System.out.println("Cost for Department " + this.name + " is " + sum);

        return sum;
    }
}

class Employee implements IDepartment {

    int costToCompany;

    Employee(int costToCompany) {
        this.costToCompany = costToCompany;
    }

    @Override
    public int getTotalCost() {
        System.out.println("Employee cost: " + costToCompany);
        return costToCompany;
    }
}

class Composite {
    public static void main(String args[]) {
        Employee e1 = new Employee(10);
        Employee e2 = new Employee(20);
        Employee e3 = new Employee(5);

        Department engineeringDepartment = new Department("engineering", new ArrayList<>(Arrays.asList(e1, e2)));
        Department hrDepartment = new Department("HR", new ArrayList<>(Arrays.asList(e3)));

        Department company = new Department("Company",
                new ArrayList<>(Arrays.asList(engineeringDepartment, hrDepartment)));
        int totalCost = company.getTotalCost();

        System.out.println(totalCost);

    }
}