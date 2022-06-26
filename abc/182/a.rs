use proconio::input;

fn main() {
    input! {
        a: i32,
        b: i32,
    }
    println!("{}", a*2 + 100 - b);   
}