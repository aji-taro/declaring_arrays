fn main() {
    println!("# Rust");

    // 基本
    let array01 = [1, 2, 3];
    println!("array01 {:?} ", array01);

    // 二次元配列
    let array02: [[i32; 3]; 2] = [[10, 20, 30], [100, 200, 300]];
    // あえてループしてみる
    for i in 0..array02.len() {
        for k in 0..array02[i].len() {
            println!("array02: {}", array02[i][k]);
        }
    }

    // 要素の追加 ※vec
    let mut vector01 = vec![11, 22, 33];
    vector01.push(44);
    println!("vector01: {:?}", &vector01);

    // 連想配列 ※vec
    let vector02 = vec![
        ("goals".to_owned(), 3),
        ("assists".to_owned(), 5),
        ("duels".to_owned(), 7),
    ];
    println!("vector02: {:?}", &vector02);
}
