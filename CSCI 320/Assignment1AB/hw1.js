function readFile(file) {
    const fs = require('fs');
    const fileData = fs.readFile(file, 'utf8');
    // split the string into an array of lines
    const lines = fileData.split('\n');
    // create an array to store the data
    const data = [];
    // iterate over the lines and split each line by the delimiter
    for (const word of lines) {
        data.push(word.split(','));
    }
    return data;
}

// bubble sort algorithm
function bubbleSortAlp(arr) {
  var temp;
  for (var i = 0; i < arr.length - 1; i++) {
      for (var j = 0; j < arr.length - i - 1; j++) {
          if (arr[j][1].localeCompare(arr[j + 1][1]) > 0) {
              temp = arr[j];
              arr[j] = arr[j + 1];
              arr[j + 1] = temp;
          }
      }
  }
  return arr;
}



// read the file
const data = readFile('Hw1_SampleFile.txt');
// sort the data
const sortedData = bubbleSortAlp(data);
//print to console
console.log(sortedData);