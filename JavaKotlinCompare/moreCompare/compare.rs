fn main() {
	let mut total = 0;
	for i in 0..99998 {
	    for j in 1..99998 {
		total += i / j;
	    }
	}
	println!("{}", total);
}
