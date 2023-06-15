import { readFileSync } from 'fs';

// Code for taking in command line arguments

const fileName : string = process.argv[2];
const fieldIndex: number = Number(process.argv[3]);

// Read the input file and split it by each line
const fileInfo: string = readFileSync(fileName, 'utf8');
const sent: string[] = fileInfo.split('\n');

// Split each line into an array of fields
const fieldsArray = sent.map(line => line.split(','));

// Merge sort function
// takes in 2 parameters, array and field and compare
function Sort(arr: string[][], field: number) {
  arr.sort((x, y) => x[field].localeCompare(y[field]));
  return arr;
}
// takes in 3 parameters, arr left/right and field #
function merge(left: string[][], right: string[][], field: number) {
  //empty array for final answer
  let final: string[][] = [], leftpos = 0, rightPos = 0;

  while (leftpos < left.length && rightPos < right.length) {
    //Compares elements at the "field"of left array at the left index postition with the right
    if (left[leftpos][field].localeCompare(right[rightPos][field]) < 0) {
      //If true, the element of the "left" array at the leftPos is pushed into the "final" array and "leftpos" index is incremented by 1.
      final.push(left[leftpos]);
      leftpos++;
    } else { //vice versa
      final.push(right[rightPos]);
      rightPos++;
    }
  }
  // The function compar elements in the left and right array and finally add them to the final array sorted. slice into new array
  return final.concat(left.slice(leftpos)).concat(right.slice(rightPos));
}

// Sort the fields array based on the specified field index
const ss = Sort(fieldsArray, fieldIndex);

// loop through and print each array element
for (let i = 0; i < ss.length; i++) {
  console.log(ss[i].join(','))
}