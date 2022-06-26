use proconio::input;

fn main() {
    input! {
        n: i32,
        x: i32,
        t: i32,
    }
    println!("{}", ((n-1)/x + 1)*t);
}