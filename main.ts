console.log('# TypeScript');

// 基本
const array01 = [1, 2, 3];
console.log('array01', array01, 'length:', array01.length);

// 二次元配列
let array02: number[][] = [[10, 20, 30], [100, 200, 300]];
// あえてループしてみる
for (let i:number=0; i<array02.length; i++) {
  for (let k:number=0; k<array02[i].length; k++) {
    console.log(`array02[${i}][${k}]`, array02[i][k]);
  }
}

// 要素の追加
let array03 = [11, 22, 33];
array03.push(44, 55);
console.log('array03', array03, 'length:', array03.length);

// 連想配列 その 1
const array04: { [key: string]: number } = {
    goals: 3,
    assists: 5,
    duels: 7,
  };
console.log('array04', array04, 'length:', Object.keys(array04).length);

// 連想配列 その 2 (type を組み合わせてみる)
type ArrayType = { [key: string]: string|number };
const array05: ArrayType = {
    id: 123,
    name: 'taro',
    age: 200,
    location: 'umi',
};
console.log('array05', array05, 'length:', Object.keys(array05).length);

