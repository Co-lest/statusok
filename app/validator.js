function validateSolution(solution) {
    try {
      const [numbers, result] = solution.split(':');
      const [num1, num2] = numbers.split(',').map(Number);
      
      if (isNaN(num1) || isNaN(num2) || isNaN(Number(result))) {
        return {
          success: false,
          message: 'Invalid input format. Expected "number1,number2:result"'
        };
      }
  
      const expectedSum = num1 + num2;
      const providedSum = Number(result);
  
      return {
        success: expectedSum === providedSum,
        message: expectedSum === providedSum ? 
          'Correct!' : 
          `Expected ${expectedSum}, but got ${providedSum}`
      };
    } catch (error) {
      return {
        success: false,
        message: 'Invalid solution format'
      };
    }
  }