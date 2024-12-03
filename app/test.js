const net = require('net');
const chalk = require('chalk');

const client = new net.Socket();

const testCases = [
  '2,3:5',    // correct
  '10,5:15',  // correct
  '1,2:4',    // incorrect
  'a,b:c',    // invalid format
];

let currentTest = 0;

client.connect(3000, 'localhost', () => {
  console.log(chalk.blue('Connected to server'));
});

client.on('data', (data) => {
  console.log(chalk.yellow('Server:', data.toString()));
  
  if (currentTest < testCases.length) {
    const test = testCases[currentTest];
    console.log(chalk.blue(`Sending test case: ${test}`));
    client.write(test + '\n');
    currentTest++;
  } else {
    client.end();
  }
});

client.on('close', () => {
  console.log(chalk.red('Connection closed'));
});