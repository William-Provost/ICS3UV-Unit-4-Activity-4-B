/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-12-07
 * @fileoverview This program shows the random built in function.
 */

// constants and variables
let someRandomFloat: number = 0.0;
let someRandomInteger: number = 0;
let someRandomIntegerBetween10And20: number = 0;

// get a random float between 0 and 1
someRandomFloat = Math.random();

// get a random integer between 0 and 100
// Math.random() generates [0, 1). Multiplying by 101 gives [0, 101). Math.floor() gives integer [0, 100].
someRandomInteger = Math.floor(Math.random() * 101);

// get a random integer between 10 and 20
// Math.random() * 11 gives [0, 11). Math.floor() gives integer [0, 10]. Adding 10 gives [10, 20].
someRandomIntegerBetween10And20 = Math.floor(Math.random() * 11) + 10;

// display results
console.log("Here are some random numbers:");
console.log(someRandomFloat);
console.log(someRandomInteger);
console.log(someRandomIntegerBetween10And20);

console.log("\nDone.");
